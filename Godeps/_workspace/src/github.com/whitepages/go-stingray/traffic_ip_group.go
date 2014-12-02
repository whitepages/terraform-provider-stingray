package stingray

import (
	"encoding/json"
	"net/http"
)

// A TrafficIPGroup is a Stingray traffic IP group.
type TrafficIPGroup struct {
	jsonResource             `json:"-"`
	TrafficIPGroupProperties `json:"properties"`
}

type TrafficIPGroupProperties struct {
	Basic struct {
		Enabled        *bool           `json:"enabled,omitempty"`
		HashSourcePort *bool           `json:"hash_source_port,omitempty"`
		IPMapping      *IPMappingTable `json:"ip_mapping,omitempty"`
		IPAddresses    *[]string       `json:"ipaddresses,omitempty"`
		KeepTogether   *bool           `json:"keeptogether,omitempty"`
		Location       *int            `json:"location,omitempty"`
		Machines       *[]string       `json:"machines,omitempty"`
		Mode           *string         `json:"mode,omitempty"`
		Multicast      *string         `json:"multicast,omitempty"`
		Note           *string         `json:"note,omitempty"`
		Slaves         *[]string       `json:"slaves,omitempty"`
	} `json:"basic"`
}

type IPMappingTable []IPMapping

type IPMapping struct {
	IP             *string `json:"ip,omitempty"`
	TrafficManager *string `json:"traffic_manager,omitempty"`
}

func (r *TrafficIPGroup) endpoint() string {
	return "traffic_ip_groups"
}

func (r *TrafficIPGroup) String() string {
	s, _ := json.Marshal(r)
	return string(s)
}

func (r *TrafficIPGroup) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewTrafficIPGroup(name string) *TrafficIPGroup {
	r := new(TrafficIPGroup)
	r.setName(name)
	return r
}

func (c *Client) GetTrafficIPGroup(name string) (*TrafficIPGroup, *http.Response, error) {
	r := NewTrafficIPGroup(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListTrafficIPGroups() ([]string, *http.Response, error) {
	return c.List(&TrafficIPGroup{})
}
