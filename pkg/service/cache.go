package service

// cacheService wraps another service, but caches it's results for the methods.
// This is a very basic cache, there is no cache eviction or timeouts. It just
// holds onto values for the lifetime of the application.
type cacheService struct {
	service        Service
	authorities    []Authority
	establishments map[string][]Establishment
}

// NewCache returns a new service that will consume a service, but acts as
// middleware for caching the results.
func NewCache(service Service) Service {
	return &cacheService{
		service:        service,
		authorities:    make([]Authority, 0),
		establishments: make(map[string][]Establishment),
	}
}

// Authorities returns a series of Authorities from the underlying API or it
// returns an error if it was not able to request or parse the result.
func (s *cacheService) Authorities() ([]Authority, error) {
	if len(s.authorities) > 0 {
		return s.authorities, nil
	}
	return s.service.Authorities()
}

// EstablishmentsForAuthority returns a series of Establishments from the
// underlying API or it returns an error if it was not able to request or
// parse the result. The Establishments service API takes a Authority
// LocalID to select the correct set of establishments for that Authority.
func (s *cacheService) EstablishmentsForAuthority(localID string) ([]Establishment, error) {
	if e, ok := s.establishments[localID]; ok && len(e) > 0 {
		return e, nil
	}
	return s.service.EstablishmentsForAuthority(localID)
}
