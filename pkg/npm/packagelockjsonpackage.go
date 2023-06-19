package npm

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type PackageLockJSONPackage struct {
	Integrity string
	Resolved  string
	Name      string
	Version   string

	Bin interface{}

	// name:constraint
	// Requires map[string]string
}

func PackageLockJSONPackageFromPVD(
	pvd *PackageVersionData,
) *PackageLockJSONPackage {

	return &PackageLockJSONPackage{
		Integrity: pvd.Dist.Integrity,
		Resolved:  pvd.Dist.Tarball,
		Name:      pvd.Name,
		Version:   pvd.Version,
		Bin:       pvd.Bin,
	}
}

func LoadPackageLockJSONPackageMetadatasFromDir(path string) ([]*PackageLockJSONPackage, error) {
	pkgs := []*PackageLockJSONPackage{}

	if err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(path, ".metadata.json") {
			metaBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			meta := &PackageLockJSONPackage{}
			if err := json.Unmarshal(metaBytes, meta); err != nil {
				return err
			}

			pkgs = append(pkgs, meta)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return pkgs, nil
}
