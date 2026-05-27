package clients

import "github.com/andrew-hayworth22/wodify-go/internal/httpclient"

// Client provides access to the Wodify Clients API.
type Client struct {
	hc *httpclient.Client
}

// New returns a new Clients client.
func New(hc *httpclient.Client) *Client {
	return &Client{hc: hc}
}
