package service

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
)

const (
	defaultHeaderTimeout = 5 * time.Second
	defaultTimeout       = 10 * time.Second
	defaultKeepAlive     = 30 * time.Second
)

// realService defines a structure for requesting entities from the ratings gov site
type realService struct {
	base    string
	version int
	client  *http.Client
	logger  log.Logger
}

// New creates a Service from a base url and the API version to use for the
// underlying service.
// Note: if a version is not supplied with the request then calls to the API
// endpoints will return no data.
func New(base string, version int, logger log.Logger) Service {
	// Create a new http client, so we can handle timeouts in a more granular
	// manor.
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			ResponseHeaderTimeout: defaultHeaderTimeout,
			Dial: (&net.Dialer{
				Timeout:   defaultTimeout,
				KeepAlive: defaultKeepAlive,
			}).Dial,
			TLSHandshakeTimeout: defaultTimeout,
			DisableKeepAlives:   false,
			MaxIdleConnsPerHost: 1,
		},
	}
	// Return the service.
	return &realService{
		base:    base,
		version: version,
		client:  client,
		logger:  logger,
	}
}

// Authorities returns a series of Authorities from the underlying API or it
// returns an error if it was not able to request or parse the result.
func (s *realService) Authorities() ([]Authority, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Authorities", s.base), nil)
	if err != nil {
		return nil, err
	}

	// Make sure we set the service API version, otherwise we get nothing.
	req.Header.Set(serviceAPIVersion, strconv.Itoa(s.version))

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Parse out the errors
	var res Authorities
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res.Authorities, nil
}

// EstablishmentsForAuthority returns a series of Establishments from the
// underlying API or it returns an error if it was not able to request or
// parse the result. The Establishments service API takes a Authority
// LocalID to select the correct set of establishments for that Authority.
func (s *realService) EstablishmentsForAuthority(id string) ([]Establishment, error) {
	return nil, nil
}
