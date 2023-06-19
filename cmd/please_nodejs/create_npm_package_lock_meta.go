package main

import (
	"encoding/json"
	"os"

	"github.com/VJftw/please-nodejs/pkg/npm"
	"github.com/urfave/cli/v2"
)

func CreateNPMPackageLockMeta() *cli.Command {
	return &cli.Command{
		Name:  "create_npm_package_lock_meta",
		Usage: "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "version",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "view",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "out",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			viewJSONBytes, err := os.ReadFile(cCtx.String("view"))
			if err != nil {
				return err
			}

			meta, err := npm.PackageLockJSONPackageFromNameVersionViewJSON(
				cCtx.String("name"),
				cCtx.String("version"),
				viewJSONBytes,
			)
			if err != nil {
				return err
			}

			metaJSONBytes, err := json.Marshal(meta)
			if err != nil {
				return err
			}

			if err := os.WriteFile(cCtx.String("out"), metaJSONBytes, 0660); err != nil {
				return err
			}

			return nil
		},
	}
}
