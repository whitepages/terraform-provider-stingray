package stingray

import "net/http"

type Rule struct {
	fileResource
}

func (r *Rule) endpoint() string {
	return "rules"
}

func NewRule(name string) *Rule {
	r := new(Rule)
	r.setName(name)
	return r
}

func (c *Client) GetRule(name string) (*Rule, *http.Response, error) {
	r := NewRule(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListRules() ([]string, *http.Response, error) {
	return c.List(&Rule{})
}
