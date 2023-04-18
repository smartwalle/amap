package amap_test

import (
	"github.com/smartwalle/amap"
	"strings"
	"testing"
)

func TestClient_Geo(t *testing.T) {
	var c = amap.New("2222879dfcb3f4bed02003140c2ce5e4")
	rsp, _ := c.Geo("", "四川省成都市金堂县测试测试。")
	t.Log(rsp.IsSuccess(), rsp.Error)
	t.Log(strings.Split(string(rsp.GeoCodes[0].Location), ","))
}

func TestClient_ReGeo(t *testing.T) {
	var c = amap.New("2222879dfcb3f4bed02003140c2ce5e4")
	//rsp, err := c.ReGeo("104.114616", "30.588039")
	rsp, err := c.ReGeo("104.056154", "30.538463")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rsp.IsSuccess(), rsp.Error)

	var info string
	if rsp.ReGeoCode != nil {

		if len(rsp.ReGeoCode.POIs) > 0 {
			info = string(rsp.ReGeoCode.POIs[0].Name)
		} else if rsp.ReGeoCode.AddressComponent != nil {
			var street = rsp.ReGeoCode.AddressComponent.StreetNumber
			if street != nil {
				info = string(street.Street + street.Number)
			}

			if info == "" {
				info = string(rsp.ReGeoCode.AddressComponent.TownShip)
			}
		}
	}

	t.Log(info)
}
