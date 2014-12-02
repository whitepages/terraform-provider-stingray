package main

import (
	"crypto/tls"
	"net/http"

	"github.com/whitepages/terraform-provider-stingray/Godeps/_workspace/src/github.com/whitepages/go-stingray"
)

// Config is the configuration structure used to instantiate the Stingray
// provider.
type Config struct {
	URL       string
	Username  string
	Password  string
	VerifySSL bool
}

func (c *Config) Client() (*stingray.Client, error) {
	client := newClient(c)

	return client, nil
}

func newClient(c *Config) *stingray.Client {
	if c.VerifySSL {
		return stingray.NewClient(nil, c.URL, c.Username, c.Password)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}

	return stingray.NewClient(httpClient, c.URL, c.Username, c.Password)
}
