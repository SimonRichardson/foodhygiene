package query

import (
	"encoding/json"
	"net/http"
)

// JSONError replies to the request with the specified error message and HTTP
// code. It does not otherwise end the request; the caller should ensure no
// further writes are done to w.
// The error message should be json compatible otherwise this will panic.
func JSONError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(code)

	// Note: it's possible that this can fail and therefore we panic.
	if e := json.NewEncoder(w).Encode(rawError{
		Error: err,
		Code:  code,
	}); e != nil {
		panic(e)
	}
}

type rawError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}
