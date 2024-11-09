package api

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, status int, reply any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if reply != nil {
		_ = json.NewEncoder(w).Encode(reply)
	}
}
