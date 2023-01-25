package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/VJftw/please-nodejs/pkg/npm/pnpm"
	"github.com/peterbourgon/mergemap"
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
			&cli.StringSliceFlag{
				Name: "merge_files",
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

			if err := depsManager.CreatePackageJSON(packageJSONOut); err != nil {
				return err
			}

			filesToMerge := cCtx.StringSlice("merge_files")
			if len(filesToMerge) < 1 {
				return nil
			}

			packageJSONContents, err := os.ReadFile(packageJSONOut)
			if err != nil {
				return err
			}
			packageJSONMapping := map[string]interface{}{}
			if err := json.Unmarshal(packageJSONContents, &packageJSONMapping); err != nil {
				return err
			}

			for _, fileToMerge := range filesToMerge {
				fileContents, err := os.ReadFile(fileToMerge)
				if err != nil {
					return err
				}

				fileMapping := map[string]interface{}{}
				if err := json.Unmarshal(fileContents, &fileMapping); err != nil {
					return err
				}

				packageJSONMapping = mergemap.Merge(fileMapping, packageJSONMapping)
			}

			mergedPackageJSONBytes, err := json.Marshal(packageJSONMapping)
			if err != nil {
				return err
			}

			return os.WriteFile(packageJSONOut, mergedPackageJSONBytes, 0644)
		},
	}
}
