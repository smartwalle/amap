package amap

import (
	"context"
	"fmt"
	"github.com/smartwalle/ngx"
)

const (
	kGeo   = "geocode/geo"
	kReGeo = "geocode/regeo"
)

// Geo 地理编码 https://lbs.amap.com/api/webservice/guide/api/georegeo#geo
func (this *Client) Geo(city, address string) (result *Geo, err error) {
	var api = fmt.Sprintf("%s/%s", this.apiDomain, kGeo)
	var req = ngx.NewRequest(ngx.Get, api, ngx.WithClient(this.Client))
	req.AddParam("key", this.key)
	req.AddParam("city", city)
	req.AddParam("address", address)

	var rsp = req.Exec(context.Background())
	if rsp.Error() != nil {
		return nil, err
	}
	if err = rsp.UnmarshalJSON(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// ReGeo 逆地理编码 https://lbs.amap.com/api/webservice/guide/api/georegeo#regeo
func (this *Client) ReGeo(longitude, latitude string) (result *ReGeo, err error) {
	var api = fmt.Sprintf("%s/%s", this.apiDomain, kReGeo)
	var req = ngx.NewRequest(ngx.Get, api, ngx.WithClient(this.Client))
	req.AddParam("key", this.key)
	req.AddParam("location", fmt.Sprintf("%s,%s", longitude, latitude))
	req.AddParam("extensions", "all")

	var rsp = req.Exec(context.Background())
	if rsp.Error() != nil {
		return nil, err
	}

	if err = rsp.UnmarshalJSON(&result); err != nil {
		return nil, err
	}
	return result, nil
}
