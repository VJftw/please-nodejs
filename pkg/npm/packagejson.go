package npm

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// PackageJSON is used to construct a package.json and package-lock.json.
type PackageJSON struct {
	Dependencies map[string]string `json:"dependencies"`
}

func NewPackageJSON() *PackageJSON {
	return &PackageJSON{
		Dependencies: map[string]string{},
	}
}

func (p *PackageJSON) SetDependency(name string, constraint string) {
	p.Dependencies[name] = constraint
}

func (p *PackageJSON) AddDependency(name string, version string) {
	p.Dependencies[name] = version
}

func (p *PackageJSON) WriteToFile(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("  ", "  ")

	if err := enc.Encode(p); err != nil {
		return err
	}

	return nil
}
