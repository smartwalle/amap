package amap4go

type GeoCode struct {
	FormattedAddress String `json:"formatted_address"`
	Country          String `json:"country"`
	Province         String `json:"province"`
	CityCode         String `json:"citycode"`
	City             String `json:"city"`
	District         String `json:"district"`
	TownShip         String `json:"township"`
	Neighborhood     struct {
		Name String `json:"name"`
		Type String `json:"type"`
	} `json:"neighborhood"`
	Building struct {
		Name String `json:"name"`
		Type String `json:"type"`
	} `json:"building"`
	AdCode   String `json:"adcode"`
	Street   String `json:"street"`
	Number   String `json:"number"`
	Location String `json:"location"`
	Level    String `json:"level"`
}

type Geo struct {
	Error
	Count    string     `json:"count"`
	GeoCodes []*GeoCode `json:"geocodes"`
}

type ReGeo struct {
	Error
	ReGeoCode *GeoCode
}
