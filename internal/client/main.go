package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/inveracity/go-embed-spa/internal/common"
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

	var hello common.Hello
	if err := json.Unmarshal(body, &hello); err != nil {
		return "", err
	}

	ret := fmt.Sprintf("Hello %s", hello.Hello)
	return ret, nil
}

func (c *Client) Memory() (string, error) {
	url := fmt.Sprintf("%s/stream/memory", c.ENDPOINT)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return "", err
		}
		l := string(line)
		if strings.Contains(l, "data:") {
			content := strings.Split(l, "data: ")
			fmt.Println(content)
		}
	}
}
