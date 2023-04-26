package amap

import "net/http"

const (
	kAMapURL = "https://restapi.amap.com/v3"
)

type Client struct {
	Client    *http.Client
	key       string
	apiDomain string
}

func New(key string) *Client {
	var c = &Client{}
	c.Client = http.DefaultClient
	c.key = key
	c.apiDomain = kAMapURL
	return c
}
