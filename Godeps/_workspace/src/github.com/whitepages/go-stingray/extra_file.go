package stingray

import "net/http"

// An ExtraFile is a Stingray extra file.
type ExtraFile struct {
	fileResource
}

func (r *ExtraFile) endpoint() string {
	return "extra_files"
}

func NewExtraFile(name string) *ExtraFile {
	r := new(ExtraFile)
	r.setName(name)
	return r
}

func (c *Client) GetExtraFile(name string) (*ExtraFile, *http.Response, error) {
	r := NewExtraFile(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListExtraFiles() ([]string, *http.Response, error) {
	return c.List(&ExtraFile{})
}
