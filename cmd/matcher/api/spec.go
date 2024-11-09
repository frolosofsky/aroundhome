package api

type Error struct {
	Message string `json:"message"`
}

type PartnerMatchResult struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Rating   string `json:"rating"`
	Distance int    `json:"distance"`
}
