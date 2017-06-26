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

	records := make([]OutputAuthority, len(r.Records))
	for k, v := range r.Records {
		records[k] = OutputAuthority{
			Name:    v.Name,
			LocalID: v.LocalID,
		}
	}

	if err := json.NewEncoder(w).Encode(records); err != nil {
		panic(err)
	}
}

// OutputAuthority is a normalized version of service.Authority. This exists
// for a couple of reasons.
// 1. The service.Authority payload is semantically confused when it comes to
//  naming. See "authorties" vs "Name" for one example.
// 2. There are some values which are not required for the UI, which helps
//  reduce the size of the payload.
type OutputAuthority struct {
	Name    string `json:"name"`
	LocalID int    `json:"local_id"`
}
