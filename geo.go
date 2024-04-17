package amap

import (
	"context"
	"net/url"
	"strings"
)

const (
	kGeo   = "v3/geocode/geo"
	kReGeo = "v3/geocode/regeo"
)

// Geo 地理编码 https://lbs.amap.com/api/webservice/guide/api/georegeo#geo
func (client *Client) Geo(ctx context.Context, city, address string) (result *Geo, err error) {
	var param = url.Values{}
	param.Add("city", city)
	param.Add("address", address)

	if err = client.request(ctx, kGeo, param, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ReGeo 逆地理编码 https://lbs.amap.com/api/webservice/guide/api/georegeo#regeo
func (client *Client) ReGeo(ctx context.Context, longitude, latitude string) (result *ReGeo, err error) {
	var param = url.Values{}
	param.Add("location", strings.Join([]string{longitude, latitude}, ","))
	param.Add("extensions", "all")

	if err = client.request(ctx, kReGeo, param, &result); err != nil {
		return nil, err
	}
	return result, nil
}
