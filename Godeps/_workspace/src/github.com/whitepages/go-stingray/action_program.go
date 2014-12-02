package stingray

import "net/http"

// An ActionProgram is Stingray action program.
type ActionProgram struct {
	fileResource
}

func (r *ActionProgram) endpoint() string {
	return "action_programs"
}

func NewActionProgram(name string) *ActionProgram {
	r := new(ActionProgram)
	r.setName(name)
	return r
}

func (c *Client) GetActionProgram(name string) (*ActionProgram, *http.Response, error) {
	r := NewActionProgram(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListActionPrograms() ([]string, *http.Response, error) {
	return c.List(&ActionProgram{})
}
