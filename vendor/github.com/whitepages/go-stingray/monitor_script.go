package stingray

import "net/http"

// A MonitorScript is a Stingray monitor program.
type MonitorScript struct {
	fileResource
}

func (r *MonitorScript) endpoint() string {
	return "monitor_scripts"
}

func NewMonitorScript(name string) *MonitorScript {
	r := new(MonitorScript)
	r.setName(name)
	return r
}

func (c *Client) GetMonitorScript(name string) (*MonitorScript, *http.Response, error) {
	r := NewMonitorScript(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListMonitorScripts() ([]string, *http.Response, error) {
	return c.List(&MonitorScript{})
}
