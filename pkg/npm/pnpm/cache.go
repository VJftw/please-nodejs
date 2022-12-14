package pnpm

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/VJftw/please-nodejs/pkg/npm"
	homedir "github.com/mitchellh/go-homedir"
)

// Cache implements methods for interacting with the `pnpm` cache.
type Cache struct {
	BaseDir string
}

// NewCache returns a new Cache with the given base dir.
func NewCache(baseDir string) (*Cache, error) {
	absBaseDir, err := homedir.Expand(baseDir)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(absBaseDir, 0774); err != nil {
		return nil, fmt.Errorf("could not create cache base dir '%s': %w", absBaseDir, err)
	}

	return &Cache{
		BaseDir: absBaseDir,
	}, nil
}

// PackageMetadata returns the NPM Package Metadata for the given pkgName, in
// the given registry from the Cache.
func (c *Cache) PackageMetadata(registry string, pkgName string) (*npm.PackageMetadata, error) {
	packageMetadataPath := filepath.Join(c.BaseDir, "metadata", registry, pkgName+".json")
	packageMetadataFile, err := os.Open(packageMetadataPath)
	if err != nil {
		return nil, fmt.Errorf("could not open package metadata file '%s': %w", packageMetadataPath, err)
	}

	packageMetadata := &npm.PackageMetadata{}
	if err := json.NewDecoder(packageMetadataFile).Decode(packageMetadata); err != nil {
		return nil, fmt.Errorf("could not decode package metadata from '%s': %w", packageMetadataPath, err)
	}

	return packageMetadata, nil
}

// PackageMetadataVersion returns the NPM Package Metadata Version for the given
// version, pkgName, in the given registry from the Cache.
func (c *Cache) PackageMetadataVersion(registry string, pkgName string, version string) (*npm.PackageVersionData, error) {
	packageMetadata, err := c.PackageMetadata(registry, pkgName)
	if err != nil {
		return nil, err
	}

	if pvd, ok := packageMetadata.Versions[version]; ok {
		return pvd, nil
	}

	return nil, fmt.Errorf("version '%s' doest not exist for '%s'", version, pkgName)
}
