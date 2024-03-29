package query

import (
	"net/http"
	"strings"
	"time"

	"github.com/SimonRichardson/foodhygiene/pkg/service"
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

// These are the query API URL paths.
const (
	APIPathAuthorities    = "/authorities"
	APIPathEstablishments = "/establishments"
)

// API serves the query API
type API struct {
	service service.Service
	logger  log.Logger
}

// NewAPI creates a API with correct dependencies.
func NewAPI(service service.Service, logger log.Logger) *API {
	return &API{
		service: service,
		logger:  logger,
	}
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	iw := &interceptingWriter{http.StatusOK, w}
	w = iw
	// Routing table
	method, path := r.Method, r.URL.Path
	switch {
	case method == "GET" && path == APIPathAuthorities:
		a.handleAuthorities(w, r)
	case method == "GET" && path == APIPathEstablishments:
		a.handleEstablishments(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (a *API) handleAuthorities(w http.ResponseWriter, r *http.Request) {
	// useful metrics
	begin := time.Now()

	defer r.Body.Close()

	// Let's guard against invalid content-types
	if !validContentType(r) {
		JSONError(w, "invalid content type", http.StatusBadRequest)
		return
	}

	// Get the authorities from the service
	authorities, err := a.service.Authorities()
	if err != nil {
		// Wrap the error request, so that we're more specific
		e := errors.Wrap(err, "error requesting authorities")
		JSONError(w, e.Error(), http.StatusInternalServerError)
		return
	}

	// AuthoritiesResult prints out the json
	qr := AuthoritiesResult{
		Duration: time.Since(begin).String(),
		Records:  authorities,
	}
	qr.EncodeTo(w)
}

func (a *API) handleEstablishments(w http.ResponseWriter, r *http.Request) {
	// useful metrics
	begin := time.Now()

	defer r.Body.Close()

	// Let's guard against invalid content-types
	if !validContentType(r) {
		JSONError(w, "invalid content type", http.StatusBadRequest)
		return
	}

	// Validate user input
	var p EstablishmentsQueryParams
	if err := p.DecodeFrom(r.URL, queryRequired); err != nil {
		JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	establishments, err := a.service.EstablishmentsForAuthority(p.LocalID)
	if err != nil {
		e := errors.Wrapf(err, "error requesting establishments for authority %q", p.LocalID)
		JSONError(w, e.Error(), http.StatusInternalServerError)
		return
	}

	// Calculate the ratings of the whole establishments for the authority
	ratings := calculateRatings(establishments)

	// EstablishmentsResult prints out the json
	qr := EstablishmentsResult{
		Params:   p,
		Duration: time.Since(begin).String(),
		Records:  ratings,
	}
	qr.EncodeTo(w)
}

// Validate the header content-type.
func validContentType(r *http.Request) bool {
	t := r.Header.Get("Content-Type")
	return strings.Contains(t, "application/json")
}

type interceptingWriter struct {
	code int
	http.ResponseWriter
}

func (iw *interceptingWriter) WriteHeader(code int) {
	iw.code = code
	iw.ResponseWriter.WriteHeader(code)
}

const (
	httpHeaderDuration = "X-Proxy-Duration"
	httpHeaderLocalID  = "X-Local-ID"
)
