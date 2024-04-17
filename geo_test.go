package amap_test

import (
	"context"
	"github.com/smartwalle/amap"
	"testing"
)

func TestClient_Geo(t *testing.T) {
	var c = amap.New("2222879dfcb3f4bed02003140c2ce5e4")
	result, err := c.Geo(context.Background(), "", "四川省成都市金堂县测试测试。")
	if err != nil {
		t.Error(err)
		return
	}
	if !result.IsSuccess() {
		t.Error(result.Error)
		return
	}

	t.Log("获得信息：", result.GeoCodes[0].Location)
}

func TestClient_ReGeo(t *testing.T) {
	var c = amap.New("2222879dfcb3f4bed02003140c2ce5e4")
	result, err := c.ReGeo(context.Background(), "104.056154", "30.538463")
	if err != nil {
		t.Error(err)
		return
	}
	if !result.IsSuccess() {
		t.Error(result.Error)
		return
	}

	var info string
	if result.ReGeoCode != nil {
		if len(result.ReGeoCode.POIs) > 0 {
			info = string(result.ReGeoCode.POIs[0].Name)
		} else if result.ReGeoCode.AddressComponent != nil {
			var street = result.ReGeoCode.AddressComponent.StreetNumber
			if street != nil {
				info = string(street.Street + street.Number)
			}

			if info == "" {
				info = string(result.ReGeoCode.AddressComponent.TownShip)
			}
		}
	}

	if info != "天府软件园G1幢" {
		t.Fatalf("期望获得：天府软件园G1幢, 实际获得：%s", info)
	}
	t.Log("获得信息：", info)
}
