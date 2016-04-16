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
