package models

type ApiStruct struct {
	Face []FaceStruct `json:"face"`
}

type FaceStruct struct {
	Attribute AttributeStruct `json:"attribute"`
}

type AttributeStruct struct {
	Age     AgeStruct     `json:"age"`
	Gender  GenderStruct  `json:"gender"`
	Smiling SmilingStruct `json:"smiling"`
}

type AgeStruct struct {
	Range int `json:"range"`
	Value int `json:"value"`
}

type GenderStruct struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
}

type SmilingStruct struct {
	Value float64 `json:"value"`
}

type GeoCode struct {
	Result []GeoResStruct `json:"results"`
}

type GeoResStruct struct {
	Formatted string `json:"formatted_address"`
}

//TextSearch

type Hotels struct {
	Results []HotelResStruct `json:"results"`
}

type HotelResStruct struct {
	Geometry GeoStruct `json:"geometry"`
	Name     string    `json:"name"`
	Rating   float64   `json:"rating"`
}

type GeoStruct struct {
	Location LocStruct `json:"location"`
}

type LocStruct struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

//JSON for getting list of places
type Places struct {
	Routes []RouteStruct `json:"routes"`
}

type RouteStruct struct {
	Legs []LegStruct `json:"legs"`
}

type LegStruct struct {
	EndAddress string `json:"end_address"`
}

//restaurant
type RestRespStruct struct {
	BreakFast []RestResponse `json:"breakfast"`
	Lunch     []RestResponse `json:"lunch"`
	Dinner    []RestResponse `json:"dinner"`
	Itenary   []IteStruct    `json:"itenary"`
}

type RestResponse struct {
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}

type IteStruct struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
