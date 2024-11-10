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

type PartnerDetails struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Rating  int      `json:"rating"`
	Skills  []string `json:"skills"`
	Address string   `json:"address"`
	Radius  int      `json:"radius"`
}
