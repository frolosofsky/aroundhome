package api

import (
	"log"
	"net/http"

	"github.com/frolosofsky/aroundhome/pkg/utils"
)

func (svc *Service) HandleMatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	query := r.URL.Query()
	material := query.Get("material")
	address := query.Get("address")

	if len(material) == 0 {
		writeResponse(w, http.StatusBadRequest, Error{`query parameter "material" is required`})
		return
	}

	if len(address) == 0 {
		writeResponse(w, http.StatusBadRequest, Error{`query parameter "address" is required`})
		return
	}

	pos, err := utils.ParsePosition(address)
	if err != nil {
		log.Printf("[debug] failed to parse address from %s: %s", address, err)
		writeResponse(w, http.StatusBadRequest, Error{`query parameter "address" must follow "longitude;latitude format"`})
		return
	}

	partners, err := svc.PartnerStore.MatchPartners(material, pos)
	if err != nil {
		log.Printf("[error] failed to match partners: %s", err)
		writeResponse(w, http.StatusInternalServerError, nil)
		return
	}

	res := make([]PartnerMatchResult, 0, len(partners))
	for _, p := range partners {
		res = append(res, PartnerMatchResult{
			Id:       p.Id,
			Name:     p.Name,
			Rating:   p.Rating,
			Distance: p.Distance,
		})
	}
	writeResponse(w, http.StatusOK, res)
}
