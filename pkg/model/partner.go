package model

type Position struct {
	Latitude  float64
	Longitude float64
}

type Partner struct {
	Id       string
	Name     string
	Position Position
	Radius   int
	Rating   int
	Skills   []string
}

type PartnerMatchResult struct {
	Id       string
	Name     string
	Position Position
	Rating   int
	Distance int
}
