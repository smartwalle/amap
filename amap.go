package amap

import (
	"context"
	"encoding/json"
	"github.com/smartwalle/ngx"
	"net/http"
	"net/url"
)

const (
	kAMap = "https://restapi.amap.com/"
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

func (client *Client) request(ctx context.Context, api string, param url.Values, result interface{}) error {
	param.Add("key", client.key)
	var req = ngx.NewRequest(ngx.Get, client.host+api, ngx.WithClient(client.Client), ngx.WithForm(param))
	var rsp, err = req.Do(ctx)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if err = json.NewDecoder(rsp.Body).Decode(&result); err != nil {
		return err
	}
	return nil
}
