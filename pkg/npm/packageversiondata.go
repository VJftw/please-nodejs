package npm

import (
	"encoding/json"
	"fmt"
)

type PackageVersionData struct {
	Name    string      `json:"name"`
	Version string      `json:"version"`
	License string      `json:"license"`
	Bin     interface{} `json:"bin"`
	Dist    struct {
		Integrity string `json:"integrity"`
		Tarball   string `json:"tarball"`
	} `json:"dist"`
	Dependencies map[string]string `json:"dependencies"`
}

type PackageVersionDataLicense string

func (pvd *PackageVersionData) UnmarshalJSON(b []byte) error {

	type PackageVersionDataAlias PackageVersionData
	aux := &struct {
		*PackageVersionDataAlias
		License interface{} `json:"license"`
	}{
		PackageVersionDataAlias: (*PackageVersionDataAlias)(pvd),
	}

	if err := json.Unmarshal(b, aux); err != nil {
		return err
	}

	switch v := aux.License.(type) {
	case string:
		pvd.License = v
	case map[string]interface{}:
		pvd.License = v["type"].(string)
	}

	return nil
}

func (pvd *PackageVersionData) String() string {
	return fmt.Sprintf("%s@%s", pvd.Name, pvd.Version)
}
