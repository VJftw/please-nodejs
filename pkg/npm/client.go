package npm

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
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
	log.Println(u.String())

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
