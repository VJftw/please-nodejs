package main

import (
	"os"
	"path/filepath"

	"github.com/VJftw/please-nodejs/pkg/npm/pnpm"
	"github.com/urfave/cli/v2"
)

func PackageJSONCommand() *cli.Command {
	return &cli.Command{
		Name:  "packagejson",
		Usage: "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "out",
				Value: "package.json",
			},
		},
		Action: func(cCtx *cli.Context) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			depsManager, err := pnpm.NewDepsManager(cwd)
			if err != nil {
				return err
			}

			packageJSONOut := cCtx.String("out")

			if !filepath.IsAbs(packageJSONOut) {
				packageJSONOut = filepath.Join(cwd, packageJSONOut)
			}

			return depsManager.CreatePackageJSON(packageJSONOut)
		},
	}
}
