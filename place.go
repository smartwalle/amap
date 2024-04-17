package amap

import (
	"context"
	"net/url"
	"strconv"
)

const (
	kPlaceTextSearch = "v3/place/text"
)

// PlaceTextSearch 搜索POI https://developer.amap.com/api/webservice/guide/api-advanced/search
func (client *Client) PlaceTextSearch(ctx context.Context, city, keywords string, page int) (result *PlaceTextSearch, err error) {
	var param = url.Values{}
	param.Add("city", city)
	param.Add("keywords", keywords)
	param.Add("children", "1")
	param.Add("extensions", "all")
	param.Add("citylimit", "true")
	param.Add("offset", "25")
	param.Add("page", strconv.Itoa(page))

	if err = client.request(ctx, kPlaceTextSearch, param, &result); err != nil {
		return nil, err
	}
	return result, nil
}
