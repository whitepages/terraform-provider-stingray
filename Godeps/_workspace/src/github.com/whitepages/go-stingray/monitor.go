package stingray

import (
	"encoding/json"
	"net/http"
)

// A Monitor is a Stingray monitor.
type Monitor struct {
	jsonResource      `json:"-"`
	MonitorProperties `json:"properties"`
}

type MonitorProperties struct {
	Basic struct {
		BackOff  *bool   `json:"back_off,omitempty"`
		Delay    *int    `json:"delay,omitempty"`
		Failures *int    `json:"failures,omitempty"`
		Machine  *string `json:"machine,omitempty"`
		Note     *string `json:"note,omitempty"`
		Scope    *string `json:"scope,omitempty"`
		Timeout  *int    `json:"timeout,omitempty"`
		Type     *string `json:"type,omitempty"`
		UseSSL   *bool   `json:"use_ssl,omitempty"`
		Verbose  *bool   `json:"verbose,omitempty"`
	} `json:"basic"`
	HTTP struct {
		Authentication *string `json:"authentication,omitempty"`
		BodyRegex      *string `json:"body_regex,omitempty"`
		HostHeader     *string `json:"host_header,omitempty"`
		Path           *string `json:"path,omitempty"`
		StatusRegex    *string `json:"status_regex,omitempty"`
	} `json:"http"`
	RTSP struct {
		BodyRegex   *string `json:"body_regex,omitempty"`
		Path        *string `json:"path,omitempty"`
		StatusRegex *string `json:"status_regex,omitempty"`
	} `json:"rtsp"`
	Script struct {
		Arguments *ScriptArgumentsTable `json:"arguments,omitempty"`
		Program   *string               `json:"program,omitempty"`
	} `json:"script"`
	SIP struct {
		BodyRegex   *string `json:"body_regex,omitempty"`
		StatusRegex *string `json:"status_regex,omitempty"`
		Transport   *string `json:"transport,omitempty"`
	} `json:"sip"`
	TCP struct {
		CloseString    *string `json:"close_string,omitempty"`
		MaxResponseLen *int    `json:"max_response_len,omitempty"`
		ResponseRegex  *string `json:"response_regex,omitempty"`
		WriteString    *string `json:"write_string,omitempty"`
	} `json:"tcp"`
	UDP struct {
		AcceptAll *bool `json:"accept_all,omitempty"`
	} `json:"udp"`
}

type ScriptArgumentsTable []ScriptArgument

type ScriptArgument struct {
	Name        *string `json:"name,omitempty"` // mandatory
	Description *string `json:"description,omitempty"`
	Value       *string `json:"value,omitempty"` // mandatory
}

func (r *Monitor) endpoint() string {
	return "monitors"
}

func (r *Monitor) String() string {
	s, _ := json.Marshal(r)
	return string(s)
}

func (r *Monitor) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewMonitor(name string) *Monitor {
	r := new(Monitor)
	r.setName(name)
	return r
}

func (c *Client) GetMonitor(name string) (*Monitor, *http.Response, error) {
	r := NewMonitor(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListMonitors() ([]string, *http.Response, error) {
	return c.List(&Monitor{})
}
