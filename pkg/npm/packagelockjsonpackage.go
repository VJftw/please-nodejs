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

func PackageLockJSONPackageFromNameVersionViewJSON(
	name string,
	version string,
	distJSONBytes []byte,
) (*PackageLockJSONPackage, error) {
	var view struct {
		Bin  interface{} `json:"bin"`
		Dist struct {
			Integrity string `json:"integrity"`
			Tarball   string `json:"tarball"`
		} `json:"dist"`
	}
	if err := json.Unmarshal(distJSONBytes, &view); err != nil {
		return nil, err
	}

	return &PackageLockJSONPackage{
		Integrity: view.Dist.Integrity,
		Resolved:  view.Dist.Tarball,
		Name:      name,
		Version:   version,
		Bin:       view.Bin,
	}, nil
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
