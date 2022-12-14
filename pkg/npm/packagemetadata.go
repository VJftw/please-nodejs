package npm

import (
	"fmt"
	"sort"

	"github.com/Masterminds/semver/v3"
)

type PackageMetadata struct {
	Name     string                         `json:"name"`
	Versions map[string]*PackageVersionData `json:"versions"`
}

func (m *PackageMetadata) GetLatestCompatibleVersionData(constraint string) (*PackageVersionData, error) {
	coll := []*semver.Version{}
	for v := range m.Versions {
		semver, err := semver.NewVersion(v)
		if err != nil {
			return nil, err
		}

		coll = append(coll, semver)
	}

	sort.Sort(sort.Reverse(semver.Collection(coll)))

	versionConstraint, err := semver.NewConstraint(constraint)
	if err != nil {
		return nil, err
	}

	for _, v := range coll {
		if versionConstraint.Check(v) {
			return m.Versions[v.Original()], nil
		}
	}

	return nil, fmt.Errorf("could not find compatible version of '%s' for '%s'", m.Name, constraint)
}
