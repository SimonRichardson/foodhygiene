package query

import (
	"net/http"

	"github.com/go-kit/kit/log"
)

// These are the query API URL paths.
const (
	APIPathAuthorities    = "/authorities"
	APIPathEstablishments = "/establishments"
)

// API serves the query API
type API struct {
	logger log.Logger
}

// NewAPI creates a API with correct dependencies.
func NewAPI(logger log.Logger) *API {
	return &API{
		logger: logger,
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

}

func (a *API) handleEstablishments(w http.ResponseWriter, r *http.Request) {

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
)
