package leads

import (
	"github.com/andrew-hayworth22/wodify-go/internal/httpclient"
)

// Client provides access to the Wodify Leads API.
type Client struct {
	hc *httpclient.Client
}

// New returns a new Leads client.
func New(hc *httpclient.Client) *Client {
	return &Client{hc: hc}
}
