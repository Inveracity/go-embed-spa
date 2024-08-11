package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/inveracity/go-embed-spa/internal/server"
)

type Client struct {
	ENDPOINT string
}

func New(ENDPOINT string) *Client {
	return &Client{ENDPOINT: ENDPOINT}
}

func (c *Client) Hello() (string, error) {
	url := fmt.Sprintf("%s/api", c.ENDPOINT)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var hello server.Hello
	if err := json.Unmarshal(body, &hello); err != nil {
		return "", err
	}

	ret := fmt.Sprintf("Hello %s", hello.Hello)
	return ret, nil
}
