package stingray

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the GitHub client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a github.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// Stingray client configured to use test server
	client = NewClient(nil, server.URL, "username", "password")
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	url := "http://localhost:9070/"
	username := "username"
	password := "password"

	c := NewClient(nil, url, username, password)

	if c.Client == nil {
		t.Errorf("NewClient Client is nil")
	}

	if c.BaseURL.String() != url {
		t.Errorf("NewClient URL = %v, want %v", c.BaseURL.String(), url)
	}

	if c.Username != username {
		t.Errorf("NewClient USERNAME = %v, want %v", c.Username, username)
	}

	if c.Password != password {
		t.Errorf("NewClient PASSWORD = %v, want %v", c.Password, password)
	}
}
