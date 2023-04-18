package amap

type Client struct {
	key       string
	apiDomain string
}

func New(key string) *Client {
	var c = &Client{}
	c.key = key
	c.apiDomain = kAPIDomain
	return c
}
