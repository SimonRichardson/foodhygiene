package query

import (
	"encoding/json"
	"net/http"

	"github.com/SimonRichardson/foodhygiene/pkg/service"
)

// AuthoritiesResult outputs the authorities from the food hygiene service
type AuthoritiesResult struct {
	Duration string
	Records  []service.Authority
}

// EncodeTo encodes the AuthoritiesResult to the HTTP response writer.
// Note: if the records can't be encoded then panic, so we don't fail silently.
func (r *AuthoritiesResult) EncodeTo(w http.ResponseWriter) {
	w.Header().Set(httpHeaderDuration, r.Duration)

	if err := json.NewEncoder(w).Encode(r.Records); err != nil {
		panic(err)
	}
}
