package query

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"reflect"

	"github.com/SimonRichardson/foodhygiene/pkg/service"
	"github.com/SimonRichardson/foodhygiene/pkg/service/mock_service"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
)

func TestAPIAuthorities(t *testing.T) {
	t.Parallel()

	t.Run("zero authorities", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mock   = mock_service.NewMockService(ctrl)
			api    = NewAPI(mock, log.NewNopLogger())
			server = httptest.NewServer(api)

			u = fmt.Sprintf("%s/authorities", server.URL)
		)
		defer server.Close()

		mock.EXPECT().
			Authorities().
			Return([]service.Authority{}, nil)

		res, err := request(u)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusOK, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}

		var auth []OutputAuthority
		if err := json.NewDecoder(res.Body).Decode(&auth); err != nil {
			t.Fatal(err)
		}
		if expected, actual := 0, len(auth); expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("one authority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mock   = mock_service.NewMockService(ctrl)
			api    = NewAPI(mock, log.NewNopLogger())
			server = httptest.NewServer(api)

			u = fmt.Sprintf("%s/authorities", server.URL)
		)
		defer server.Close()

		name, localID := "Yorkshire", 123

		mock.EXPECT().
			Authorities().
			Return([]service.Authority{
				service.Authority{
					Name:    name,
					LocalID: localID,
				},
			}, nil)

		res, err := request(u)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusOK, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}

		var auth []OutputAuthority
		if err := json.NewDecoder(res.Body).Decode(&auth); err != nil {
			t.Fatal(err)
		}
		if expected, actual := 1, len(auth); expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
		if expected, actual := (OutputAuthority{name, localID}), auth[0]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mock   = mock_service.NewMockService(ctrl)
			api    = NewAPI(mock, log.NewNopLogger())
			server = httptest.NewServer(api)

			u = fmt.Sprintf("%s/authorities", server.URL)
		)
		defer server.Close()

		mock.EXPECT().
			Authorities().
			Return(nil, errors.New("something went wrong"))

		res, err := request(u)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusInternalServerError, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestAPIEstablishments(t *testing.T) {
	t.Parallel()

	t.Run("zero rating", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mock   = mock_service.NewMockService(ctrl)
			api    = NewAPI(mock, log.NewNopLogger())
			server = httptest.NewServer(api)

			u = fmt.Sprintf("%s/establishments?local_id=0", server.URL)
		)
		defer server.Close()

		mock.EXPECT().
			EstablishmentsForAuthority("0").
			Return([]service.Establishment{}, nil)

		res, err := request(u)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusOK, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}

		var rate []OutputRating
		if err := json.NewDecoder(res.Body).Decode(&rate); err != nil {
			t.Fatal(err)
		}
		if expected, actual := 0, len(rate); expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("one rating", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mock   = mock_service.NewMockService(ctrl)
			api    = NewAPI(mock, log.NewNopLogger())
			server = httptest.NewServer(api)

			u = fmt.Sprintf("%s/establishments?local_id=0", server.URL)
		)
		defer server.Close()

		name, rating := "Bobs burgers", "4"

		mock.EXPECT().
			EstablishmentsForAuthority("0").
			Return([]service.Establishment{
				service.Establishment{
					Name:   name,
					Rating: rating,
				},
			}, nil)

		res, err := request(u)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusOK, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}

		var rate []OutputRating
		if err := json.NewDecoder(res.Body).Decode(&rate); err != nil {
			t.Fatal(err)
		}
		if expected, actual := 1, len(rate); expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
		if expected, actual := (OutputRating{fmt.Sprintf("%s-Star", rating), "100.00%"}), rate[0]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mock   = mock_service.NewMockService(ctrl)
			api    = NewAPI(mock, log.NewNopLogger())
			server = httptest.NewServer(api)

			u = fmt.Sprintf("%s/establishments?local_id=0", server.URL)
		)
		defer server.Close()

		mock.EXPECT().
			EstablishmentsForAuthority("0").
			Return(nil, errors.New("something went wrong"))

		res, err := request(u)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusInternalServerError, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("error no local id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mock   = mock_service.NewMockService(ctrl)
			api    = NewAPI(mock, log.NewNopLogger())
			server = httptest.NewServer(api)

			u = fmt.Sprintf("%s/establishments", server.URL)
		)
		defer server.Close()

		res, err := request(u)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusBadRequest, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func request(u string) (*http.Response, error) {
	req, err := http.NewRequest("GET", u, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	return client.Do(req)
}
