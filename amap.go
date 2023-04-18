package amap

const (
	kAMapURL = "https://restapi.amap.com/v3"
)

type Client struct {
	key       string
	apiDomain string
}

func New(key string) *Client {
	var c = &Client{}
	c.key = key
	c.apiDomain = kAMapURL
	return c
}
