package stingray

import (
	"encoding/json"
	"net/http"
)

// A TrafficManager is a Stingray traffic manager.
type TrafficManager struct {
	jsonResource             `json:"-"`
	TrafficManagerProperties `json:"properties"`
}

type TrafficManagerProperties struct{}

func (r *TrafficManager) endpoint() string {
	return "traffic_managers"
}

func (r *TrafficManager) String() string {
	s, _ := json.Marshal(r)
	return string(s)
}

func (r *TrafficManager) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewTrafficManager(name string) *TrafficManager {
	r := new(TrafficManager)
	r.setName(name)
	return r
}

func (c *Client) GetTrafficManager(name string) (*TrafficManager, *http.Response, error) {
	r := NewTrafficManager(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListTrafficManagers() ([]string, *http.Response, error) {
	return c.List(&TrafficManager{})
}
