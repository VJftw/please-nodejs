package pnpm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/VJftw/please-nodejs/pkg/npm"
	"github.com/rs/zerolog/log"
)

// DepsManager implements a dependency build system compatible dependency
// manager.
type DepsManager struct {
	basePath string

	depVersions map[string]map[string]struct{}
}

func NewDepsManager(basePath string) (*DepsManager, error) {
	m := &DepsManager{
		basePath:    basePath,
		depVersions: map[string]map[string]struct{}{},
	}

	if err := m.loadDeps(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *DepsManager) MeetsDepConstraint(depName string, depConstraint string) error {
	depVersions, ok := m.depVersions[depName]
	if !ok {
		return fmt.Errorf("missing dep '%s: %s'", depName, depConstraint)
	}

	versionConstraint, err := semver.NewConstraint(depConstraint)
	if err != nil {
		return err
	}

	for depVersion := range depVersions {
		semver, err := semver.NewVersion(depVersion)
		if err != nil {
			return err
		}

		if versionConstraint.Check(semver) {
			return nil
		}
	}

	return fmt.Errorf("could not find compatible version for '%s: %s' in: %v", depName, depConstraint, depVersions)
}

func (m *DepsManager) CreatePackageJSON(path string) error {
	pj := npm.NewPackageJSON()

	for depName, depVersions := range m.depVersions {
		for depVersion := range depVersions {
			pj.SetDependency(depName, depVersion)
		}
	}

	return pj.WriteToFile(path)
}

func (m *DepsManager) loadDeps() error {
	if err := filepath.Walk(
		m.basePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			if !info.IsDir() && filepath.Ext(path) == ".tgz" {
				// most likely an npm pkg
				pkgName, pkgVersion, err := m.parseDepNameAndVersionFromTgz(path)
				if err != nil {
					return err
				}
				if _, ok := m.depVersions[pkgName]; !ok {
					m.depVersions[pkgName] = map[string]struct{}{}
				}

				m.depVersions[pkgName][pkgVersion] = struct{}{}
				log.Debug().
					Str("name", pkgName).
					Str("version", pkgVersion).
					Str("path", path).
					Msg("found npm package")
			}

			return nil
		},
	); err != nil {
		return fmt.Errorf("could not load deps from '%s': %w", m.basePath, err)
	}

	return nil
}

func (m *DepsManager) parseDepNameAndVersionFromTgz(path string) (string, string, error) {
	fileName := filepath.Base(path)
	cleanPkgNameVersion := strings.TrimSuffix(fileName, ".tgz")
	// npm package scopes are fixed at 1 level of nesting.
	pkgNameVersion := strings.Replace(cleanPkgNameVersion, "*", "/", 1)

	pkgNameVersionParts := strings.Split(pkgNameVersion, "@")

	switch len(pkgNameVersionParts) {
	case 2:
		return pkgNameVersionParts[0], pkgNameVersionParts[1], nil
	case 3:
		return "@" + pkgNameVersionParts[1], pkgNameVersionParts[2], nil
	}

	return "", "", fmt.Errorf("invalid pkg name version: %s", pkgNameVersion)
}
