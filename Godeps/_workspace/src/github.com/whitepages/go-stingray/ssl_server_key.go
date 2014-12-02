package stingray

import (
	"encoding/json"
	"net/http"
)

// A SSLServerKey is a Stingray trusted certificate.
type SSLServerKey struct {
	jsonResource           `json:"-"`
	SSLServerKeyProperties `json:"properties"`
}

type SSLServerKeyProperties struct {
	Basic struct {
		Note    *string `json:"note,omitempty"`
		Private *string `json:"private,omitempty"`
		Public  *string `json:"public,omitempty"`
		Request *string `json:"request,omitempty"`
	} `json:"basic"`
}

func (r *SSLServerKey) endpoint() string {
	return "ssl/server_keys"
}

func (r *SSLServerKey) String() string {
	s, _ := json.Marshal(r)
	return string(s)
}

func (r *SSLServerKey) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewSSLServerKey(name string) *SSLServerKey {
	r := new(SSLServerKey)
	r.setName(name)
	return r
}

func (c *Client) GetSSLServerKey(name string) (*SSLServerKey, *http.Response, error) {
	r := NewSSLServerKey(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListSSLServerKey() ([]string, *http.Response, error) {
	return c.List(&SSLServerKey{})
}
