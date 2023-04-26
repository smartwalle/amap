package amap

import "net/http"

const (
	kAMap = "https://restapi.amap.com/v3"
)

type Client struct {
	Client *http.Client
	key    string
	host   string
}

func New(key string) *Client {
	var c = &Client{}
	c.Client = http.DefaultClient
	c.key = key
	c.host = kAMap
	return c
}
