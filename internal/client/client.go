package client

import (
	"fmt"
	"io"
	"net/http"
)

type client struct {
	client *http.Client
}

func NewClient() *client {
	return &client{
		client: http.DefaultClient,
	}
}

func (c *client) RequestServer(url string) error {
	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	fmt.Println(string(data))

	return nil
}
