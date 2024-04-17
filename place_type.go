package amap

type PlaceTextSearch struct {
	Error
	Count      String     `json:"count"`
	Suggestion Suggestion `json:"suggestion"`
	POIs       []*POI     `json:"pois"`
}

type Suggestion struct {
	Keywords []String `json:"keywords"`
	Cities   []*City  `json:"cities"`
}

type City struct {
	Name String `json:"name"`
	//Num      String `json:"num"`
	CityCode String `json:"citycode"`
	AdCode   String `json:"adcode"`
}
