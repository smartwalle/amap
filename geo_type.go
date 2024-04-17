package amap

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
	Id           String        `json:"id"`
	Parent       String        `json:"parent"`
	ChildType    String        `json:"childtype"`
	Name         String        `json:"name"`
	Tag          String        `json:"tag"`
	Type         String        `json:"type"`
	TypeCode     String        `json:"typecode"`
	BizType      String        `json:"biz_type"`
	Address      String        `json:"address"`
	Location     String        `json:"location"`
	Tel          String        `json:"tel"`
	Postcode     String        `json:"postcode"`
	Website      String        `json:"website"`
	Email        String        `json:"email"`
	PCode        String        `json:"pcode"`
	PName        String        `json:"pname"`
	CityCode     String        `json:"citycode"`
	CityName     String        `json:"cityname"`
	AdCode       String        `json:"adcode"`
	AdName       String        `json:"adname"`
	Importance   String        `json:"importance"`
	ShopId       String        `json:"shopid"`
	ShopInfo     String        `json:"shopinfo"`
	POIWeight    String        `json:"poiweight"`
	GridCode     String        `json:"gridcode"`
	Distance     String        `json:"distance"`
	NaviPOIId    String        `json:"navi_poiid"`
	EntrLocation String        `json:"entr_location"`
	BusinessArea String        `json:"business_area"`
	ExitLocation String        `json:"exit_location"`
	Match        String        `json:"match"`
	Recommend    String        `json:"recommend"`
	Timestamp    String        `json:"timestamp"`
	Alias        String        `json:"alias"`
	IndoorMap    String        `json:"indoor_map"`
	IndoorData   Indoor        `json:"indoor_data"`
	GroupBuyNum  String        `json:"groupbuy_num"`
	DiscountNum  String        `json:"discount_num"`
	BizExt       BizExt        `json:"biz_ext"`
	Event        []interface{} `json:"event"`
	Children     []interface{} `json:"children"`
	Photos       []Photo       `json:"photos"`
}

type Indoor struct {
	Cpid      String `json:"cpid"`
	Floor     String `json:"floor"`
	Truefloor String `json:"truefloor"`
	Cmsid     String `json:"cmsid"`
}

type BizExt struct {
	Rating String `json:"rating"`
	Cost   String `json:"cost"`
}

type Photo struct {
	Title String `json:"title"`
	URL   String `json:"url"`
}

type ReGeo struct {
	Error
	ReGeoCode *ReGeoCode `json:"regeocode"`
}
