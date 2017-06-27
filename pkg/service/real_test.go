package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestRealServiceAuthorities(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		var (
			api     = http.NewServeMux()
			server  = httptest.NewServer(api)
			service = New(server.URL, 2, log.NewNopLogger())
		)
		defer server.Close()

		api.HandleFunc("/Authorities", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			res := Authorities{
				Authorities: []Authority{},
			}
			if err := json.NewEncoder(w).Encode(res); err != nil {
				t.Fatal(err)
			}
		})

		got, err := service.Authorities()
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := 0, len(got); expected != actual {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("one", func(t *testing.T) {
		var (
			api     = http.NewServeMux()
			server  = httptest.NewServer(api)
			service = New(server.URL, 2, log.NewNopLogger())
		)
		defer server.Close()

		auth := Authority{
			Name:    "Yorkshire",
			LocalID: 123,
		}

		api.HandleFunc("/Authorities", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			res := Authorities{
				Authorities: []Authority{
					auth,
				},
			}
			if err := json.NewEncoder(w).Encode(res); err != nil {
				t.Fatal(err)
			}
		})

		got, err := service.Authorities()
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := 1, len(got); expected != actual {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
		if expected, actual := auth, got[0]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("error", func(t *testing.T) {
		var (
			api     = http.NewServeMux()
			server  = httptest.NewServer(api)
			service = New(server.URL, 2, log.NewNopLogger())
		)
		defer server.Close()

		api.HandleFunc("/Authorities", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		_, err := service.Authorities()
		if expected, actual := true, err != nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestRealServiceEstablishmentsForAuthority(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		var (
			api     = http.NewServeMux()
			server  = httptest.NewServer(api)
			service = New(server.URL, 2, log.NewNopLogger())
		)
		defer server.Close()

		api.HandleFunc("/Establishments", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			res := Establishments{
				Establishments: []Establishment{},
			}
			if err := json.NewEncoder(w).Encode(res); err != nil {
				t.Fatal(err)
			}
		})

		got, err := service.EstablishmentsForAuthority("0")
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := 0, len(got); expected != actual {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("one", func(t *testing.T) {
		var (
			api     = http.NewServeMux()
			server  = httptest.NewServer(api)
			service = New(server.URL, 2, log.NewNopLogger())
		)
		defer server.Close()

		est := Establishment{
			Name:   "Bobs Burgers",
			Rating: "4",
		}

		api.HandleFunc("/Establishments", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			res := Establishments{
				Establishments: []Establishment{
					est,
				},
			}
			if err := json.NewEncoder(w).Encode(res); err != nil {
				t.Fatal(err)
			}
		})

		got, err := service.EstablishmentsForAuthority("0")
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := 1, len(got); expected != actual {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
		if expected, actual := est, got[0]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("error", func(t *testing.T) {
		var (
			api     = http.NewServeMux()
			server  = httptest.NewServer(api)
			service = New(server.URL, 2, log.NewNopLogger())
		)
		defer server.Close()

		api.HandleFunc("/Establishments", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		_, err := service.EstablishmentsForAuthority("0")
		if expected, actual := true, err != nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}
