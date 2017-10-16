// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "nmgapi": Image Resource Client
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
	"strconv"
)

// ShowImagePath computes a request path to the show action of Image.
func ShowImagePath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/nmg/image/%s", param0)
}

// Show an image's metadata
func (c *Client) ShowImage(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowImageRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowImageRequest create the request corresponding to the show action endpoint of the Image resource.
func (c *Client) NewShowImageRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UploadImagePath computes a request path to the upload action of Image.
func UploadImagePath(id string) string {
	param0 := id

	return fmt.Sprintf("/nmg/image/%s", param0)
}

// Upload multiple images in multipart request
func (c *Client) UploadImage(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewUploadImageRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUploadImageRequest create the request corresponding to the upload action endpoint of the Image resource.
func (c *Client) NewUploadImageRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}