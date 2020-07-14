package amap4go

type GeoCode struct {
	FormattedAddress String        `json:"formatted_address"`
	Country          String        `json:"country"`
	Province         String        `json:"province"`
	CityCode         String        `json:"citycode"`
	City             String        `json:"city"`
	District         String        `json:"district"`
	TownShip         String        `json:"township"`
	Neighborhood     *Neighborhood `json:"neighborhood"`
	Building         *Building     `json:"building"`
	AdCode           String        `json:"adcode"`
	Street           String        `json:"street"`
	Number           String        `json:"number"`
	Location         String        `json:"location"`
	Level            String        `json:"level"`
}

type Building struct {
	Name String `json:"name"`
	Type String `json:"type"`
}

type Neighborhood struct {
	Name String `json:"name"`
	Type String `json:"type"`
}

type Geo struct {
	Error
	Count    string     `json:"count"`
	GeoCodes []*GeoCode `json:"geocodes"`
}

type ReGeoCode struct {
	Roads            []*Road           `json:"roads"`
	AddressComponent *AddressComponent `json:"addressComponent"`
	FormattedAddress String            `json:"formatted_address"`
	POIs             []*POI            `json:"pois"`
}

type AddressComponent struct {
	City         String        `json:"city"`
	Province     String        `json:"province"`
	AdCode       String        `json:"adcode"`
	District     String        `json:"district"`
	TownCode     String        `json:"towncode"`
	StreetNumber *StreetNumber `json:"streetNumber"`
	Country      String        `json:"country"`
	TownShip     String        `json:"township"`
	//BusinessAreas []*BusinessArea `json:"businessAreas"`
	Neighborhood *Neighborhood `json:"neighborhood"`
	Building     *Building     `json:"building"`
	CityCode     String        `json:"citycode"`
}

type StreetNumber struct {
	Number    String `json:"number"`
	Location  String `json:"location"`
	Direction String `json:"direction"`
	Distance  String `json:"distance"`
	Street    String `json:"street"`
}

type BusinessArea struct {
	Id           String `json:"id"`
	Name         String `json:"name"`
	Location     String `json:"location"`
	BusinessArea String `json:"businessArea"`
}

type Road struct {
	Id           String `json:"id"`
	Location     String `json:"location"`
	Direction    String `json:"direction"`
	Name         String `json:"name"`
	BusinessArea String `json:"businessArea"`
	Distance     String `json:"distance"`
}

type POI struct {
	Id        String `json:"id"`
	Direction String `json:"direction"`
	//BusinessAreas []*BusinessArea `json:"businessAreas"`
	Address   String `json:"address"`
	PoiWeight String `json:"poiweight"`
	Name      String `json:"name"`
	Location  String `json:"location"`
	Distance  String `json:"distance"`
	Tel       String `json:"tel"`
	Type      String `json:"type"`
}

type ReGeo struct {
	Error
	ReGeoCode *ReGeoCode
}
