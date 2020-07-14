package amap4go

import (
	"strings"
	"testing"
)

func TestClient_Geo(t *testing.T) {
	var c = New("2222879dfcb3f4bed02003140c2ce5e4")
	rsp, _ := c.Geo("", "四川省成都市金堂县测试测试。")
	t.Log(rsp.IsSuccess(), rsp.Error)
	t.Log(strings.Split(string(rsp.GeoCodes[0].Location), ","))
}

func TestClient_ReGeo(t *testing.T) {
	var c = New("2222879dfcb3f4bed02003140c2ce5e4")
	rsp, _ := c.ReGeo("104.056267", "30.538432")
	t.Log(rsp.IsSuccess(), rsp.Error)
	t.Log(rsp.ReGeoCode.FormattedAddress)
}
