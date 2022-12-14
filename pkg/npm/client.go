package npm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"sort"

	"github.com/Masterminds/semver/v3"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
	cache      map[string]interface{}
}

func NewClient(rawBaseURL string) (*Client, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("could not parse base url '%s': %w", rawBaseURL, err)
	}

	return &Client{
		BaseURL:   baseURL,
		UserAgent: "",
		httpClient: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   http.DefaultClient.Timeout,
		},
	}, nil
}

func (c *Client) Package(name string) (*PackageMetadata, error) {
	rel := &url.URL{Path: path.Join(name)}
	u := c.BaseURL.ResolveReference(rel)

	if pm, ok := c.cache[u.String()]; ok {
		return pm.(*PackageMetadata), nil
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pm := &PackageMetadata{}
	if err := json.NewDecoder(resp.Body).Decode(pm); err != nil {
		return nil, fmt.Errorf("could not decode response from '%s' into %T: %w", u.String(), pm, err)
	}

	return pm, err
}

func (c *Client) PackageVersion(name string, version string) (*PackageVersionData, error) {
	rel := &url.URL{Path: path.Join(name, version)}
	u := c.BaseURL.ResolveReference(rel)

	if pvd, ok := c.cache[u.String()]; ok {
		return pvd.(*PackageVersionData), nil
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pvd := &PackageVersionData{}
	if err := json.NewDecoder(resp.Body).Decode(pvd); err != nil {
		return nil, fmt.Errorf("could not decode response from '%s' into %T: %w", u.String(), pvd, err)
	}

	return pvd, err
}

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

type PackageVersionData struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	License      string            `json:"license"`
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
