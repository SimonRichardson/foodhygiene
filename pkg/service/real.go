package service

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
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
}

// New creates a Service from a base url and the API version to use for the
// underlying service.
// Note: if a version is not supplied with the request then calls to the API
// endpoints will return no data.
func New(base string, version int) Service {
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
	return &realService{base, version, client}
}

// Authority returns a list of Authority from the underlying API service.
// Note: an error is returned if a parse error is encountered.
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
	var res []Authority
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}
