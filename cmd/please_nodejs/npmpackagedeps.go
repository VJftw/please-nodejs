package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

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
				Name:     "required_json",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			requiredDependenciesJSON, err := os.ReadFile(cCtx.String("required_json"))
			if err != nil {
				return err
			}

			requiredDependencies := map[string]string{}
			if len(strings.TrimSpace(string(requiredDependenciesJSON))) > 0 {
				if err := json.Unmarshal(requiredDependenciesJSON, &requiredDependencies); err != nil {
					return err
				}
			}

			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			availablePackages, err := npm.LoadPackageLockJSONPackageMetadatasFromDir(cwd)
			if err != nil {
				return err
			}

			// availableDependenciesTxt, err := os.ReadFile(cCtx.String("available_txt"))
			// if err != nil {
			// 	return err
			// }
			availableDependencies := map[string][]string{}
			for _, pkg := range availablePackages {
				if _, ok := availableDependencies[pkg.Name]; !ok {
					availableDependencies[pkg.Name] = []string{}
				}

				availableDependencies[pkg.Name] = append(availableDependencies[pkg.Name], pkg.Version)
			}
			// for _, line := range strings.Split(string(availableDependenciesTxt), "\n") {
			// 	if strings.TrimSpace(line) == "" {
			// 		continue
			// 	}
			// 	lineParts := strings.Split(line, "@")
			// 	var name string
			// 	var version string
			// 	switch len(lineParts) {
			// 	case 2:
			// 		name = lineParts[0]
			// 		version = lineParts[1]
			// 	case 3:
			// 		name = lineParts[0] + "@" + lineParts[1]
			// 		version = lineParts[2]
			// 	default:
			// 		return fmt.Errorf("unexpected number of @s in package name: %s", line)
			// 	}

			// 	if _, ok := availableDependencies[name]; !ok {
			// 		availableDependencies[name] = []string{}
			// 	}

			// 	availableDependencies[name] = append(availableDependencies[name], version)
			// }

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

			return nil
		},
	}
}
