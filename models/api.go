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
	Results []HotelResStruct
}

type HotelResStruct struct {
	Name     string
	Rating   float64
	Geometry GeoStruct
}

type GeoStruct struct {
	Location LocStruct
}

type LocStruct struct {
	Lat float64
	Lng float64
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
