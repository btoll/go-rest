// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "nmgapi": Enum Resource Client
//
// Command:
// $ goagen
// --design=github.com/btoll/rest-go/design
// --out=$(GOPATH)/src/github.com/btoll/rest-go
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListEnumPath computes a request path to the list action of Enum.
func ListEnumPath() string {

	return fmt.Sprintf("/nmg/enum/")
}

// Get a map of all enums
func (c *Client) ListEnum(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListEnumRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListEnumRequest create the request corresponding to the list action endpoint of the Enum resource.
func (c *Client) NewListEnumRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}