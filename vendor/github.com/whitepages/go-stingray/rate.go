package stingray

import (
	"encoding/json"
	"net/http"
)

// A Rate is a Stingray rate shaping class.
type Rate struct {
	jsonResource   `json:"-"`
	RateProperties `json:"properties"`
}

type RateProperties struct {
	Basic struct {
		MaxRatePerMinute *int    `json:"max_rate_per_minute,omitempty"`
		MaxRatePerSecond *int    `json:"max_rate_per_second,omitempty"`
		Note             *string `json:"note,omitempty"`
	} `json:"basic"`
}

func (r *Rate) endpoint() string {
	return "rate"
}

func (r *Rate) String() string {
	s, _ := jsonMarshal(r)
	return string(s)
}

func (r *Rate) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewRate(name string) *Rate {
	r := new(Rate)
	r.setName(name)
	return r
}

func (c *Client) GetRate(name string) (*Rate, *http.Response, error) {
	r := NewRate(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListRates() ([]string, *http.Response, error) {
	return c.List(&Rate{})
}
