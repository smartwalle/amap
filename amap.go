package amap4go

import "net/http"

type Client struct {
	key       string
	apiDomain string
	client    *http.Client
}

func New(key string) *Client {
	var c = &Client{}
	c.key = key
	c.apiDomain = kAPIDomain
	c.client = http.DefaultClient
	return c
}
