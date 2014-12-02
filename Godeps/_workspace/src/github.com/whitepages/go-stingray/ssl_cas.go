package stingray

import "net/http"

// An SSLCAs is a Stingray trusted certificate.
type SSLCAs struct {
	fileResource
}

func (r *SSLCAs) endpoint() string {
	return "ssl/cas"
}

func NewSSLCAs(name string) *SSLCAs {
	r := new(SSLCAs)
	r.setName(name)
	return r
}

func (c *Client) GetSSLCAs(name string) (*SSLCAs, *http.Response, error) {
	r := NewSSLCAs(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListSSLCAs() ([]string, *http.Response, error) {
	return c.List(&SSLCAs{})
}
