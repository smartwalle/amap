package amap

import (
	"context"
	"net/url"
	"strconv"
)

const (
	kDistrictSearch = "v3/config/district"
)

// DistrictSearch 行政区域查询 https://developer.amap.com/api/webservice/guide/api/district
func (client *Client) DistrictSearch(ctx context.Context, keywords, subDistrict string, page int) (result *DistrictSearch, err error) {
	var param = url.Values{}
	param.Add("keywords", keywords)
	param.Add("subdistrict", subDistrict)
	param.Add("extensions", "base")
	param.Add("offset", "25")
	param.Add("page", strconv.Itoa(page))

	if err = client.request(ctx, kDistrictSearch, param, &result); err != nil {
		return nil, err
	}
	return result, nil
}
