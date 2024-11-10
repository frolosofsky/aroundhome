package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (svc *Service) HandlePartnerDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, ok := strings.CutPrefix(r.URL.Path, "/partners/")
	if len(id) == 0 || !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	partner, err := svc.PartnerStore.GetPartner(id)
	if err != nil {
		log.Printf("[error] failed to get partern=%s details: %s", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if partner == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res := PartnerDetails{
		Id:      partner.Id,
		Name:    partner.Name,
		Rating:  partner.Rating,
		Skills:  partner.Skills,
		Address: fmt.Sprintf("%.5f,%.5f", partner.Position.Latitude, partner.Position.Longitude),
		Radius:  partner.Radius,
	}

	writeResponse(w, http.StatusOK, res)
}
