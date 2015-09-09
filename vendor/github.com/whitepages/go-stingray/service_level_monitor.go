package stingray

import (
	"encoding/json"
	"net/http"
)

// A ServiceLevelMonitor is a Stingray service level monitor.
type ServiceLevelMonitor struct {
	jsonResource                  `json:"-"`
	ServiceLevelMonitorProperties `json:"properties"`
}

type ServiceLevelMonitorProperties struct {
	Basic struct {
		Note             *string `json:"note,omitempty"`
		ResponseTime     *int    `json:"response_time,omitempty"`
		SeriousThreshold *int    `json:"serious_threshold,omitempty"`
		WarningThreshold *int    `json:"warning_threshold,omitempty"`
	} `json:"basic"`
}

func (r *ServiceLevelMonitor) endpoint() string {
	return "service_level_monitors"
}

func (r *ServiceLevelMonitor) String() string {
	s, _ := jsonMarshal(r)
	return string(s)
}

func (r *ServiceLevelMonitor) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewServiceLevelMonitor(name string) *ServiceLevelMonitor {
	r := new(ServiceLevelMonitor)
	r.setName(name)
	return r
}

func (c *Client) GetServiceLevelMonitor(name string) (*ServiceLevelMonitor, *http.Response, error) {
	r := NewServiceLevelMonitor(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListServiceLevelMonitors() ([]string, *http.Response, error) {
	return c.List(&ServiceLevelMonitor{})
}
