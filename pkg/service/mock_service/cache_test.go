package mock_service

import (
	"reflect"
	"testing"

	"github.com/SimonRichardson/foodhygiene/pkg/service"
	"github.com/golang/mock/gomock"
)

// CacheService is found with in this package, because of the way go treats
// circular dependencies i.e. it doesn't :sigh:
func TestCacheServiceAuthorities(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mock = NewMockService(ctrl)
			api  = service.NewCache(mock)
		)

		mock.EXPECT().
			Authorities().
			Return([]service.Authority{}, nil)

		_, err := api.Authorities()
		if expected, actual := true, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("repeated", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mock = NewMockService(ctrl)
			api  = service.NewCache(mock)
			auth = service.Authority{
				Name:    "Yorkshire",
				LocalID: 123,
			}
		)

		mock.EXPECT().
			Authorities().
			Return([]service.Authority{auth}, nil)

		_, err := api.Authorities()
		if expected, actual := true, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}

		// This should use the cache and not the mock.
		_, err = api.Authorities()
		if expected, actual := true, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestCacheServiceEstablishmentsForAuthority(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mock = NewMockService(ctrl)
			api  = service.NewCache(mock)
		)

		mock.EXPECT().
			EstablishmentsForAuthority("0").
			Return([]service.Establishment{}, nil)

		_, err := api.EstablishmentsForAuthority("0")
		if expected, actual := true, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("repeated", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mock = NewMockService(ctrl)
			api  = service.NewCache(mock)
			est  = service.Establishment{
				Name:   "Bobs Burgers",
				Rating: "3",
			}
		)

		mock.EXPECT().
			EstablishmentsForAuthority("0").
			Return([]service.Establishment{est}, nil)

		got, err := api.EstablishmentsForAuthority("0")
		if expected, actual := true, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
		if expected, actual := est, got[0]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}

		// This should use the cache and not the mock.
		got, err = api.EstablishmentsForAuthority("0")
		if expected, actual := true, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
		if expected, actual := est, got[0]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("vary repeated", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mock = NewMockService(ctrl)
			api  = service.NewCache(mock)
			est0 = service.Establishment{
				Name:   "Bobs Burgers",
				Rating: "3",
			}
			est1 = service.Establishment{
				Name:   "Petes Pizza",
				Rating: "4",
			}
		)

		mock.EXPECT().
			EstablishmentsForAuthority("0").
			Return([]service.Establishment{est0}, nil)

		mock.EXPECT().
			EstablishmentsForAuthority("1").
			Return([]service.Establishment{est1}, nil)

		got, err := api.EstablishmentsForAuthority("0")
		if expected, actual := true, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
		if expected, actual := est0, got[0]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}

		// This should use the cache and not the mock.
		got, err = api.EstablishmentsForAuthority("1")
		if expected, actual := true, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
		if expected, actual := est1, got[0]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}
