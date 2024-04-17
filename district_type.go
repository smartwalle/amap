package amap

type DistrictSearch struct {
	Error
	Count     string      `json:"count"`
	Districts []*District `json:"districts"`
}

type District struct {
	CityCode  String      `json:"citycode"`
	AdCode    String      `json:"adcode"`
	Name      String      `json:"name"`
	Polyline  String      `json:"polyline"`
	Center    String      `json:"center"`
	Level     String      `json:"level"`
	Districts []*District `json:"districts"`
}
