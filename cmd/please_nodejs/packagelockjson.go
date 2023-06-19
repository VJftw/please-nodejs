package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/VJftw/please-nodejs/pkg/npm"
	"github.com/urfave/cli/v2"
)

func PackageLockJSONCommand() *cli.Command {
	return &cli.Command{
		Name:  "packagelockjson",
		Usage: "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "out",
				Value: "package-lock.json",
			},
		},
		Action: func(cCtx *cli.Context) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			packages, err := npm.LoadPackageLockJSONPackageMetadatasFromDir(cwd)
			if err != nil {
				return err
			}

			type packageLockJSONDependency struct {
				Version   string `json:"version"`
				Resolved  string `json:"resolved"`
				Integrity string `json:"integrity"`
			}

			type packageLockJSONPackage struct {
				Version   string      `json:"version"`
				Resolved  string      `json:"resolved"`
				Integrity string      `json:"integrity"`
				Bin       interface{} `json:"bin,omitempty"`
			}

			packageLockJSON := &struct {
				Name            string                                `json:"name"`
				LockfileVersion int                                   `json:"lockfileVersion"`
				Requires        bool                                  `json:"requires"`
				Packages        map[string]*packageLockJSONPackage    `json:"packages"`
				Dependencies    map[string]*packageLockJSONDependency `json:"dependencies"`
			}{
				Name:            "plz-nodejs",
				LockfileVersion: 2,
				Requires:        true,
				Packages:        map[string]*packageLockJSONPackage{},
				Dependencies:    map[string]*packageLockJSONDependency{},
			}

			for _, pkg := range packages {

				packageLockJSON.Packages["node_modules/"+pkg.Name] = &packageLockJSONPackage{
					Version:   pkg.Version,
					Resolved:  pkg.Resolved,
					Integrity: pkg.Integrity,
					Bin:       pkg.Bin,
				}

				packageLockJSON.Dependencies[pkg.Name] = &packageLockJSONDependency{
					Version:   pkg.Version,
					Resolved:  pkg.Resolved,
					Integrity: pkg.Integrity,
				}

			}

			packageLockJSONBytes, err := json.Marshal(packageLockJSON)
			if err != nil {
				return err
			}

			packageLockJSONOut := cCtx.String("out")
			if !filepath.IsAbs(packageLockJSONOut) {
				packageLockJSONOut = filepath.Join(cwd, packageLockJSONOut)
			}

			if err := os.WriteFile(packageLockJSONOut, packageLockJSONBytes, 0660); err != nil {
				return err
			}

			return nil
		},
	}
}
