package service

const (
	serviceAPIVersion = "X-API-Version"
)

// Service describes a service that talks to the underlying API
// The service is envisioned as a interface so that it's possible to abstract
// the API for mocking during testing.
type Service interface {
	// Authorities returns a series of Authorities from the underlying API or it
	// returns an error if it was not able to request or parse the result.
	Authorities() ([]Authority, error)

	// EstablishmentsForAuthority returns a series of Establishments from the
	// underlying API or it returns an error if it was not able to request or
	// parse the result. The Establishments service API takes a Authority
	// LocalID to select the correct set of establishments for that Authority.
	EstablishmentsForAuthority(string) ([]Establishment, error)
}

// Authority defines a schema for the JSON from the service
type Authority struct {
	Name               string `json:"Name"`
	LocalID            int    `json:"LocalAuthorityId"`
	EstablishmentCount int    `json:"EstablishmentCount"`
}

// Establishment defines a schema for the JSON from the service
type Establishment struct {
	Name   string `json:"BusinessName"`
	Rating string `json:"RatingValue"`
}
