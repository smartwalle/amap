package amap_test

import (
	"context"
	"github.com/smartwalle/amap"
	"testing"
)

func TestClient_PlaceTextSearch(t *testing.T) {
	var c = amap.New("2222879dfcb3f4bed02003140c2ce5e4")
	result, err := c.PlaceTextSearch(context.Background(), "成都", "天府广场", 1)
	if err != nil {
		t.Error(err)
		return
	}
	if !result.IsSuccess() {
		t.Error(result.Error)
		return
	}

	if result.POIs[0].Name != "天府广场" {
		t.Fatalf("期望获得：天府广场, 实际获得：%s", result.POIs[0].Name)
	}
	t.Log("获得信息：", result.POIs[0].Name)
}
