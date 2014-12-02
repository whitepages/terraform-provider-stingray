package stingray

import "net/http"

// A LicenseKey is a Stingray license key.
type LicenseKey struct {
	fileResource
}

func (r *LicenseKey) endpoint() string {
	return "license_keys"
}

func NewLicenseKey(name string) *LicenseKey {
	r := new(LicenseKey)
	r.setName(name)
	return r
}

func (c *Client) GetLicenseKey(name string) (*LicenseKey, *http.Response, error) {
	r := NewLicenseKey(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListLicenseKeys() ([]string, *http.Response, error) {
	return c.List(&LicenseKey{})
}
