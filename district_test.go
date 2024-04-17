package amap_test

import (
	"context"
	"github.com/smartwalle/amap"
	"testing"
)

func TestClient_DistrictSearch(t *testing.T) {
	var c = amap.New("2222879dfcb3f4bed02003140c2ce5e4")
	result, err := c.DistrictSearch(context.Background(), "四川", "1", 1)
	if err != nil {
		t.Error(err)
		return
	}
	if !result.IsSuccess() {
		t.Error(result.Error)
		return
	}

	if result.Districts[0].Name != "四川省" {
		t.Fatalf("期望获得：四川省, 实际获得：%s", result.Districts[0].Name)
	}
	t.Log("获得信息：", result.Districts[0].Name)
}
