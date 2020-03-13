package mcversions

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func New(opts ...OptFunc) (*Client, error) {
	c := &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}

	for _, opt := range opts {
		if err := opt.Apply(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) GetManifest() (*http.Response, *Manifest, error) {
	return c.getManifest(ManifestURL)
}

func (c *Client) GetVersion(id string) (*http.Response, *Version, error) {
	_, man, err := c.GetManifest()
	if err != nil {
		return nil, nil, err
	}

	for _, v := range man.Versions {
		if string(v.ID) == id {
			return c.getVersion(v.URL)
		}
	}

	return nil, nil, fmt.Errorf("version id %s not found in version manifest", id)
}

func (c *Client) GetLatestRelease() (*Version, error) {
	id, err := c.GetLatestReleaseID()
	if err != nil {
		return nil, err
	}

	_, version, err := c.GetVersion(id)
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (c *Client) GetLatestSnapshot() (*Version, error) {
	id, err := c.GetLatestSnapshotID()
	if err != nil {
		return nil, err
	}

	_, version, err := c.GetVersion(id)
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (c *Client) GetLatestReleaseID() (string, error) {
	_, man, err := c.GetManifest()
	if err != nil {
		return "", err
	}

	return man.Latest.Release, nil
}

func (c *Client) GetLatestSnapshotID() (string, error) {
	_, man, err := c.GetManifest()
	if err != nil {
		return "", err
	}

	return man.Latest.Snapshot, nil
}

func (c *Client) download(url string, path string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func (c *Client) DownloadServerJar(v *Version, path string) error {
	url := v.Downloads.Server.URL
	path = filepath.Join(path, filepath.Base(url))
	return c.download(url, path)
}

func (c *Client) getManifest(url string) (*http.Response, *Manifest, error) {
	var manifest Manifest

	req := c.newRequest("GET", url, nil)

	resp, err := req.do(&manifest)
	if err != nil {
		return nil, nil, err
	}

	return resp, &manifest, nil
}

func (c *Client) getVersion(url string) (*http.Response, *Version, error) {
	var version Version

	req := c.newRequest("GET", url, nil)

	resp, err := req.do(&version)
	if err != nil {
		return nil, nil, err
	}

	return resp, &version, nil
}

func (c *Client) newRequest(method string, path string, body io.Reader) *request {
	req, err := http.NewRequest(method, path, body)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return &request{req, c, err}
}

type request struct {
	*http.Request

	client *Client
	err    error
}

func (r *request) do(v interface{}) (*http.Response, error) {
	if r.err != nil {
		return nil, r.err
	}

	resp, err := r.client.httpClient.Do(r.Request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if v != nil {
		return resp, json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, nil
}

type Opt interface {
	Apply(c *Client) error
}

type OptFunc func(c *Client) error

func (o OptFunc) Apply(c *Client) error {
	return o(c)
}

func SetClient(client *http.Client) Opt {
	return OptFunc(func(c *Client) error {
		if client == nil {
			return errors.New("client is nil")
		}
		c.httpClient = client
		return nil
	})
}
