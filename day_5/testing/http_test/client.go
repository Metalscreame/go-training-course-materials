package iterable

import (
	"fmt"
	"net/http"
)

// Client is a structure that represents client to work with Iterable service for a single client.
type Client struct {
	client *http.Client
}

func (c *Client) GetData() (err error) {
	resp, err := c.client.Get("https://google.com")
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got %v status code", resp.StatusCode)
	}

	return
}


