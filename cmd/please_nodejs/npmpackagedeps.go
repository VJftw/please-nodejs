package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Masterminds/semver/v3"
	"github.com/VJftw/please-nodejs/pkg/npm"
	"github.com/urfave/cli/v2"
)

func NPMPackageDepsCommand() *cli.Command {
	return &cli.Command{
		Name:  "npmpackagedeps",
		Usage: "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "registry",
				Value: "https://registry.npmjs.org",
			},
			&cli.StringFlag{
				Name:     "package",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "version",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "meta_out",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			npmClient, err := npm.NewClient(cCtx.String("registry"))
			if err != nil {
				return fmt.Errorf("could not initialise npm client: %w", err)
			}

			pvd, err := npmClient.PackageVersion(
				cCtx.String("package"),
				cCtx.String("version"),
			)
			if err != nil {
				return err
			}

			requiredDependencies := pvd.Dependencies

			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			availablePackages, err := npm.LoadPackageLockJSONPackageMetadatasFromDir(cwd)
			if err != nil {
				return err
			}

			availableDependencies := map[string][]string{}
			for _, pkg := range availablePackages {
				if _, ok := availableDependencies[pkg.Name]; !ok {
					availableDependencies[pkg.Name] = []string{}
				}

				availableDependencies[pkg.Name] = append(availableDependencies[pkg.Name], pkg.Version)
			}

			for rName, rConstraint := range requiredDependencies {
				versions, ok := availableDependencies[rName]
				if !ok {
					return fmt.Errorf("%s is missing", rName)
				}

				c, err := semver.NewConstraint(rConstraint)
				if err != nil {
					return err
				}

				compatible := false
				for _, version := range versions {
					v, err := semver.NewVersion(version)
					if err != nil {
						return err
					}

					if c.Check(v) {
						compatible = true
					}
				}

				if !compatible {
					return fmt.Errorf("%s@%v does not meet constraint %s: %s", rName, versions, rName, rConstraint)
				}

			}

			meta := npm.PackageLockJSONPackageFromPVD(pvd)
			metaJSONBytes, err := json.Marshal(meta)
			if err != nil {
				return err
			}

			if err := os.WriteFile(cCtx.String("meta_out"), metaJSONBytes, 0660); err != nil {
				return err
			}

			return nil
		},
	}
}
