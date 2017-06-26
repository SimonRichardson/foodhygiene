package query

import (
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"testing"
	"testing/quick"
)

func TestEstablishmentsQueryParams(t *testing.T) {
	t.Parallel()

	t.Run("decode", func(t *testing.T) {
		fn := func(a ASCII) bool {
			var (
				qp EstablishmentsQueryParams

				id     = a.String()
				u, err = url.Parse(fmt.Sprintf("http://example.com?local_id=%s", id))
			)
			if err != nil {
				t.Error(err)
			}
			if err := qp.DecodeFrom(u, queryRequired); err != nil {
				t.Error(err)
			}

			return qp.LocalID == id
		}

		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("decode required", func(t *testing.T) {
		var (
			qp     EstablishmentsQueryParams
			u, err = url.Parse("http://example.com")
		)
		if err != nil {
			t.Error(err)
		}
		if err := qp.DecodeFrom(u, queryRequired); err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("decode optional", func(t *testing.T) {
		var (
			qp     EstablishmentsQueryParams
			u, err = url.Parse("http://example.com")
		)
		if err != nil {
			t.Error(err)
		}
		if err := qp.DecodeFrom(u, queryOptional); err != nil {
			t.Errorf("expected error")
		}
	})
}

// ASCII a better string implementation for quick checking urls.
type ASCII string

// Generate allows ASCII to be used within quickcheck scenarios.
func (ASCII) Generate(r *rand.Rand, size int) reflect.Value {
	var (
		chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		res   []byte
	)

	for i := 0; i < size; i++ {
		pos := r.Intn(len(chars) - 1)
		res = append(res, chars[pos])
	}

	return reflect.ValueOf(ASCII(string(res)))
}

func (a ASCII) String() string {
	return string(a)
}
