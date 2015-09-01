package stingray

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	basePath = "/api/tm/3.5/config/active/"
)

// A Client manages communication with the Stingray API.
type Client struct {
	// HTTP client used to communicate with the API.
	Client *http.Client

	// API base URL
	BaseURL *url.URL

	// Username used for communicating with the API.
	Username string

	// Password used for communicating with the API.
	Password string
}

// NewClient returns a new Stingray API client, using the supplied
// URL, username, and password
func NewClient(httpClient *http.Client, urlStr, username string, password string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	u, _ := url.Parse(urlStr)
	if u.Path == "" {
		rel, _ := url.Parse(basePath)
		u = u.ResolveReference(rel)
	}

	c := &Client{
		Client:   httpClient,
		BaseURL:  u,
		Username: username,
		Password: password,
	}

	return c
}

// NewRequest creates a new request with the params
func (c *Client) NewRequest(method, urlStr string, body string) (*http.Request, error) {
	rel, err := url.Parse(c.BaseURL.Path + urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := strings.NewReader(body)
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	return req, nil
}

// Do sends an API request.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	return resp, nil
}

// Get retrieves a resource
func (c *Client) Get(r Resourcer) (*http.Response, error) {
	u := fmt.Sprintf("%v/%v", r.endpoint(), r.Name())

	req, err := c.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return resp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}

	err = r.decode(body)

	return resp, err
}

// Set sets a resource
func (c *Client) Set(r Resourcer) (*http.Response, error) {
	u := fmt.Sprintf("%v/%v", r.endpoint(), r.Name())

	req, err := c.NewRequest("PUT", u, r.String())
	req.Header.Add("Content-Type", r.contentType())
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)

	return resp, err
}

// Delete deletes a resource
func (c *Client) Delete(r Resourcer) (*http.Response, error) {
	u := fmt.Sprintf("%v/%v", r.endpoint(), r.Name())

	req, err := c.NewRequest("DELETE", u, "")
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)

	return resp, err
}

// List lists resources of the specified type
func (c *Client) List(r Resourcer) ([]string, *http.Response, error) {
	req, err := c.NewRequest("GET", r.endpoint(), "")
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, resp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	rl := &resourceList{}
	err = rl.decode(body)
	if err != nil {
		return nil, resp, err
	}

	return rl.names(), resp, nil
}

// ErrorResponse represents an error message returned by the Stingray API.
//
// See Chapter 2, Further Aspects of the Resource Model, Errors.
type ErrorResponse struct {
	Response  *http.Response // HTTP response that caused this error
	ID        string         `json:"error_id"`
	Text      string         `json:"error_text"`
	ErrorInfo interface{}    `json:"error_info"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.ID, r.Text, r.ErrorInfo)
}

// CheckResponse checks the API response for errors, and returns them
// if present. A response is considered an error if it has a status
// code outside the 200 range. API error responses are expected to
// have either no response body, or a JSON response body that maps to
// ErrorResponse. Any other response body will be silently ignored.
func CheckResponse(resp *http.Response) error {
	if c := resp.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: resp}
	data, err := ioutil.ReadAll(resp.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	return &v
}

// Int is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it, but unlike Int32
// its argument value is an int.
func Int(v int) *int {
	return &v
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	return &v
}

// jsonMarshal un-escapes certain "\uXXXX" escape sequences since the
// Stingray REST API does not decode these correctly. The
// json.Unmarshal function creates these escape sequences for &, <,
// and >.
func jsonMarshal(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)

	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)

	return b, err
}
