package client

import (
	"bytes"
	"encoding/json"
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

func (c *client) Request(method string, url string, query any) ([]byte, error) {
	buf, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
