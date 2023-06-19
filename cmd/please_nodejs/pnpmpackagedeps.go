package main

import (
	"fmt"
	"os"

	"github.com/VJftw/please-nodejs/pkg/npm/pnpm"
	"github.com/urfave/cli/v2"
)

func PNPMPackageDepsCommand() *cli.Command {
	return &cli.Command{
		Name:  "pnpmpackagedeps",
		Usage: "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "registry",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "version",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "cache_dir",
				Value: "~/.cache/pnpm",
			},
		},
		Action: func(cCtx *cli.Context) error {
			cache, err := pnpm.NewCache(cCtx.String("cache_dir"))
			if err != nil {
				return err
			}

			pvd, err := cache.PackageMetadataVersion(
				cCtx.String("registry"),
				cCtx.String("name"),
				cCtx.String("version"),
			)
			if err != nil {
				return err
			}

			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			depsManager, err := pnpm.NewDepsManager(cwd)
			if err != nil {
				return err
			}

			for depName, depConstraint := range pvd.Dependencies {
				if err := depsManager.MeetsDepConstraint(depName, depConstraint); err != nil {
					return fmt.Errorf("did not meet constraint in '%s': %w", pvd.Name, err)
				}
			}

			return nil
		},
	}
}
