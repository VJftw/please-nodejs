package nodejs

import (
	"strings"

	"github.com/VJftw/please-nodejs/pkg/npm"
)

// InstallableNPMPackage represents an NPM Package with each dependency version
// resolved.
type InstallableNPMPackage struct {
	*npm.PackageVersionData

	ResolvedDepVersions map[string]*npm.PackageVersionData
}

func ParseNPMPackageToNameAndVersion(npmPackage string) (string, string) {
	name := ""
	version := ""
	if npmPackage[0] == '@' {
		// we're using a scoped pkg
		name = "@"
		npmPackage = npmPackage[1:]
	}

	parts := strings.Split(npmPackage, "@")
	switch len(parts) {
	case 1:
		name += parts[0]
		version = "latest"
	case 2:
		name += parts[0]
		version = parts[1]
	}

	return name, version
}
