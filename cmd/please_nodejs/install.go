package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/VJftw/please-nodejs/pkg/npm"
	"github.com/VJftw/please-nodejs/pkg/please"
	"github.com/VJftw/please-nodejs/pkg/please/nodejs"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func InstallCommand() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Aliases: []string{"add"},
		Usage:   "Generates third party Please rules for the given NPM module",
		Description: strings.TrimSpace(`
This command generates third party Please rules for the given NPM module.
It will write third party Please rule configuration in a structured format of
BUILD files. e.g.:
  - ` + "`" + `//third_party/nodejs/react:react` + "`" + `
  - ` + "`" + `//third_party/nodejs/@babel/core:core` + "`" + `
`),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "registry",
				Value: "https://registry.npmjs.org",
			},
			&cli.IntFlag{
				Name:  "workers",
				Value: runtime.NumCPU() - 1,
			},
			&cli.IntFlag{
				Name:  "buffer_length",
				Value: 256,
				Usage: "The size of the resolver buffers. If you find this tool hanging, increase this.",
			},
			&cli.StringFlag{
				Name:  "pkg_prefix",
				Value: "//third_party/nodejs",
			},
			&cli.StringFlag{
				Name:  "nodejs_build_defs",
				Value: "///nodejs//build/defs:nodejs",
			},
			&cli.BoolFlag{
				Name:    "structured",
				Aliases: []string{"s"},
				Value:   false,
			},
		},
		Action: func(cCtx *cli.Context) error {
			npmClient, err := npm.NewClient(cCtx.String("registry"))
			if err != nil {
				return fmt.Errorf("could not initialise npm client: %w", err)
			}

			npmPackageResolver := nodejs.NewNPMPackageResolver(&nodejs.NPMPackageResolverOpts{
				Workers:       cCtx.Int("workers"),
				BufferLength:  cCtx.Int("buffer_length"),
				Toolchain:     cCtx.String("toolchain"),
				Registry:      npmClient.BaseURL.Host,
				Structured:    cCtx.Bool("structured"),
				PackagePrefix: cCtx.String("pkg_prefix"),
			}, npmClient)

			npmRules := []*please.Rule{}
			for _, requestedPkg := range cCtx.Args().Slice() {
				log.Info().Msgf("resolving npm packages for: %s", requestedPkg)
				newRules, err := npmPackageResolver.ResolveNPMPackages(cCtx.Context, requestedPkg)
				if err != nil {
					return err
				}
				npmRules = append(npmRules, newRules...)
			}

			bfm := please.NewBuildFileManager()

			for _, npmRule := range npmRules {
				bf, err := bfm.GetBuildFileForTarget(npmRule.Target)
				if err != nil {
					return err
				}

				if err := bf.UpsertRule(npmRule); err != nil {
					return err
				}

				if err := bf.EnsureSubinclude(cCtx.String("nodejs_build_defs")); err != nil {
					return err
				}
			}

			return bfm.SaveAll()
		},
	}
}
