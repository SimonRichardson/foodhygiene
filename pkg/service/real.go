package service

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

// These timeouts are required as some local authorities take some time to load,
// so tweaking these can prevent a timeout on the client side.
const (
	defaultHeaderTimeout = 10 * time.Second
	defaultTimeout       = 25 * time.Second
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
	req, err := s.newRequest("/Authorities")
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if code := resp.StatusCode; code < 200 || code >= 300 {
		return nil, errors.Errorf("invalid request (status code: %d)", code)
	}

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
	req, err := s.newRequest(fmt.Sprintf("/Establishments?localAuthorityId=%s&pageSize=0", id))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if code := resp.StatusCode; code < 200 || code >= 300 {
		return nil, errors.Errorf("invalid request (status code: %d)", code)
	}

	// Parse out the errors
	var res Establishments
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res.Establishments, nil
}

// newRequest makes sure that every request we send to the service has the
// valid headers.
func (s *realService) newRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", s.base, url), nil)
	if err != nil {
		return nil, err
	}

	// Make sure we set the service API version, otherwise we get nothing.
	req.Header.Set(serviceAPIVersion, strconv.Itoa(s.version))
	req.Header.Set(serviceContentType, contentType)

	return req, err
}
