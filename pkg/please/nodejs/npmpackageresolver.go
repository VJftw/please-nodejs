package nodejs

import (
	"context"
	"fmt"
	"math"
	"runtime"
	"strings"
	"sync"

	"github.com/VJftw/please-nodejs/pkg/npm"
	"github.com/VJftw/please-nodejs/pkg/please"
	"github.com/rs/zerolog/log"
)

type NPMPackageResolverOpts struct {
	Workers            int
	BufferLength       int
	Structured         bool
	npmPackageTargetFn func(npmPackage *npm.PackageVersionData) string
	Toolchain          string
	Registry           string
	PackagePrefix      string
}

func (r *NPMPackageResolverOpts) StructuredNPMPackageTargetFn(npmPackage *npm.PackageVersionData) string {
	return fmt.Sprintf("//%s:%s", npmPackage.Name, npmPackage.Version)
}

func (r *NPMPackageResolverOpts) SingleNPMPackageTargetFn(npmPackage *npm.PackageVersionData) string {
	return fmt.Sprintf("//:%s_%s", strings.ReplaceAll(npmPackage.Name, "/", "_"), npmPackage.Version)
}

// NPMPackageResolver implements the procedures for resolving NPM packages from
// an NPM package registry.
type NPMPackageResolver struct {
	npmClient *npm.Client

	opts *NPMPackageResolverOpts

	resolvedPackages sync.Map
}

// NewNPMPackageResolver returns a new NPMPackageResolver.
func NewNPMPackageResolver(
	opts *NPMPackageResolverOpts,
	npmClient *npm.Client,
) *NPMPackageResolver {

	if opts.Workers == 0 {
		opts.Workers = int(math.Ceil(float64(runtime.NumCPU()) / 2))
	}

	if opts.Structured {
		opts.npmPackageTargetFn = opts.StructuredNPMPackageTargetFn
	} else {
		opts.npmPackageTargetFn = opts.SingleNPMPackageTargetFn
	}

	return &NPMPackageResolver{
		npmClient:        npmClient,
		opts:             opts,
		resolvedPackages: sync.Map{},
	}
}

func (r *NPMPackageResolver) newRuleFromInstallableNPMPackage(
	npmPkg *InstallableNPMPackage,
) (*please.Rule, error) {
	npmPackageRef := r.opts.npmPackageTargetFn(npmPkg.PackageVersionData)
	target, err := please.NewTargetFromReference(npmPackageRef)
	target.PrefixPkg(r.opts.PackagePrefix)
	if err != nil {
		return nil, err
	}

	deps := []string{}
	for _, resolvedDep := range npmPkg.ResolvedDepVersions {
		depRef := r.opts.npmPackageTargetFn(resolvedDep)
		depTarget, err := please.NewTargetFromReference(depRef)
		if err != nil {
			return nil, err
		}
		depTarget.PrefixPkg(r.opts.PackagePrefix)

		if target.Pkg == depTarget.Pkg {
			// short-hand deps that are in the same pkg.
			deps = append(deps, ":"+depTarget.Name)
		} else {
			deps = append(deps, depTarget.String())
		}
	}

	rule, err := please.NewRule(
		target,
		"nodejs_npm_package",
		map[string]interface{}{
			"deps":         deps,
			"license":      npmPkg.License,
			"package_name": npmPkg.Name,
			"registry":     r.opts.Registry,
			"toolchain":    r.opts.Toolchain,
			"version":      npmPkg.Version,
			"visibility":   []string{"PUBLIC"},
		})
	if err != nil {
		return nil, err
	}

	return rule, nil
}

// parseNPMPackage fetches a single NPM package's metadata from the NPM
// registry.
func (g *NPMPackageResolver) parseNPMPackage(npmPackage string) (*npm.PackageVersionData, error) {
	name, version := ParseNPMPackageToNameAndVersion(npmPackage)

	packageVersionData, err := g.npmClient.PackageVersion(name, version)
	if err != nil {
		return nil, err
	}

	return packageVersionData, nil
}

func (r *NPMPackageResolver) ResolveNPMPackages(ctx context.Context, npmPackage string) ([]*please.Rule, error) {
	pkgsCh := make(chan string, r.opts.BufferLength)
	errsCh := make(chan error, r.opts.BufferLength)
	resCh := make(chan *InstallableNPMPackage, r.opts.BufferLength)
	resultsCh := make(chan []*please.Rule, 1)
	errCh := make(chan error, 1)

	wg := &sync.WaitGroup{}

	pkgsCh <- npmPackage
	wg.Add(1)

	for i := 0; i < r.opts.Workers; i++ {
		go r.resolveWorker(ctx, i, wg, pkgsCh, resCh, errsCh)
	}

	go func(resCh chan *InstallableNPMPackage) {
		results := []*please.Rule{}
		for res := range resCh {
			rule, err := r.newRuleFromInstallableNPMPackage(res)
			if err != nil {
				log.Error().Err(err).Msg("ERROR")
			}
			results = append(results, rule)
		}
		resultsCh <- results
	}(resCh)

	go func(errsCh chan error) {
		var errs error
		for err := range errsCh {
			if errs == nil {
				errs = err
			} else {
				errs = fmt.Errorf("%w: %w", errs, err)
			}
		}
		errCh <- errs
	}(errsCh)

	log.Debug().
		Msg("waiting for items to finish processing")
	wg.Wait()
	close(resCh)
	close(errsCh)
	defer close(resultsCh)
	defer close(errCh)

	log.Debug().
		Msg("receiving results")
	results := <-resultsCh
	log.Debug().
		Msg("receiving errs")
	err := <-errCh

	log.Debug().
		Msg("done")
	return results, err
}

func (r *NPMPackageResolver) resolveWorker(
	ctx context.Context,
	workerNum int,
	wg *sync.WaitGroup,
	pkgsCh chan string,
	resCh chan *InstallableNPMPackage,
	errCh chan error,
) {
	defer log.Debug().
		Int("workerNum", workerNum).
		Msg("stopping worker")

	log.Debug().
		Int("workerNum", workerNum).
		Msg("starting worker")

	for {
		select {
		case <-ctx.Done():
			return
		case pkg := <-pkgsCh:
			if err := r.resolveWorkerPkg(ctx, wg, pkg, pkgsCh, resCh); err != nil {
				errCh <- err
			}
		}
	}
}

func (r *NPMPackageResolver) resolveWorkerPkg(
	ctx context.Context,
	wg *sync.WaitGroup,
	pkg string,
	pkgsCh chan string,
	resCh chan *InstallableNPMPackage,
) error {
	defer wg.Done()

	pvd, err := r.parseNPMPackage(pkg)
	if err != nil {
		return err
	}

	installableNPMPackage := &InstallableNPMPackage{
		PackageVersionData:  pvd,
		ResolvedDepVersions: map[string]*npm.PackageVersionData{},
	}

	_, ok := r.resolvedPackages.Load(pvd.String())
	if ok {
		return nil
	}

	for dependencyName, versionConstraintStr := range pvd.Dependencies {
		pkgMetadata, err := r.npmClient.Package(dependencyName)
		if err != nil {
			return err
		}

		compatibleVersion, err := pkgMetadata.GetLatestCompatibleVersionData(versionConstraintStr)
		if err != nil {
			return err
		}

		_, ok := r.resolvedPackages.Load(compatibleVersion.String())
		if !ok {
			log.Debug().
				Str("package", compatibleVersion.Name).
				Str("version", compatibleVersion.Version).
				Str("pkgsCh", fmt.Sprintf("%d/%d", len(pkgsCh), cap(pkgsCh))).
				Msg("queued pkg")

			pkgsCh <- fmt.Sprintf("%s@%s", compatibleVersion.Name, compatibleVersion.Version)
			wg.Add(1)
		}

		// log.Debug().
		// 	Str("package", installableNPMPackage.Name).
		// 	Str("dep", pkgMetadata.Name).
		// 	Str("version", compatibleVersion.Version).
		// 	Msg("pkg depends on")

		installableNPMPackage.ResolvedDepVersions[pkgMetadata.Name] = compatibleVersion
	}

	r.resolvedPackages.Store(pvd.String(), installableNPMPackage)

	log.Debug().
		Str("package", installableNPMPackage.Name).
		Str("version", installableNPMPackage.Version).
		Str("resCh", fmt.Sprintf("%d/%d", len(pkgsCh), cap(pkgsCh))).
		Msg("added pkg")
	resCh <- installableNPMPackage

	return nil
}
