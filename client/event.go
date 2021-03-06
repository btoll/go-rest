// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "nmgapi": Event Resource Client
//
// Command:
// $ goagen
// --design=github.com/btoll/rest-go/design
// --out=$(GOPATH)/src/github.com/btoll/rest-go
// --version=v1.3.0

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateEventPath computes a request path to the create action of Event.
func CreateEventPath() string {

	return fmt.Sprintf("/admin/event/")
}

// Create a new sports event.
func (c *Client) CreateEvent(ctx context.Context, path string, payload *EventPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateEventRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateEventRequest create the request corresponding to the create action endpoint of the Event resource.
func (c *Client) NewCreateEventRequest(ctx context.Context, path string, payload *EventPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteEventPath computes a request path to the delete action of Event.
func DeleteEventPath(id string) string {
	param0 := id

	return fmt.Sprintf("/admin/event/%s", param0)
}

// Delete a sports event by id.
func (c *Client) DeleteEvent(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteEventRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteEventRequest create the request corresponding to the delete action endpoint of the Event resource.
func (c *Client) NewDeleteEventRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListEventPath computes a request path to the list action of Event.
func ListEventPath() string {

	return fmt.Sprintf("/admin/event/list")
}

// Get all events
func (c *Client) ListEvent(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListEventRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListEventRequest create the request corresponding to the list action endpoint of the Event resource.
func (c *Client) NewListEventRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowEventPath computes a request path to the show action of Event.
func ShowEventPath(id string) string {
	param0 := id

	return fmt.Sprintf("/admin/event/%s", param0)
}

// Get a sports event by id.
func (c *Client) ShowEvent(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowEventRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowEventRequest create the request corresponding to the show action endpoint of the Event resource.
func (c *Client) NewShowEventRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateEventPath computes a request path to the update action of Event.
func UpdateEventPath(id string) string {
	param0 := id

	return fmt.Sprintf("/admin/event/%s", param0)
}

// Update a sports event by id.
func (c *Client) UpdateEvent(ctx context.Context, path string, payload *EventPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateEventRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateEventRequest create the request corresponding to the update action endpoint of the Event resource.
func (c *Client) NewUpdateEventRequest(ctx context.Context, path string, payload *EventPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
