package query

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"

	"testing/quick"

	"github.com/SimonRichardson/foodhygiene/pkg/service"
)

func TestCalculateRatings(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		got := calculateRatings(make([]service.Establishment, 0))
		if expected, actual := 0, len(got); expected != actual {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("one", func(t *testing.T) {
		estab := []service.Establishment{
			service.Establishment{
				Name:   "Bobs burgers",
				Rating: "3",
			},
		}
		want := []Rating{
			Rating{
				Name:   "3-Star",
				Rating: 100.0,
			},
		}
		got := calculateRatings(estab)
		if expected, actual := 1, len(got); expected != actual {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
		if expected, actual := want, got; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("two same rating", func(t *testing.T) {
		estab := []service.Establishment{
			service.Establishment{
				Name:   "Bobs burgers",
				Rating: "3",
			},
			service.Establishment{
				Name:   "Freds Pizzas",
				Rating: "3",
			},
		}
		want := []Rating{
			Rating{
				Name:   "3-Star",
				Rating: 100.0,
			},
		}
		got := calculateRatings(estab)
		if expected, actual := 1, len(got); expected != actual {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
		if expected, actual := want, got; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("two different rating", func(t *testing.T) {
		estab := []service.Establishment{
			service.Establishment{
				Name:   "Bobs burgers",
				Rating: "3",
			},
			service.Establishment{
				Name:   "Freds Pizzas",
				Rating: "4",
			},
		}
		want := []Rating{
			Rating{
				Name:   "3-Star",
				Rating: 50.0,
			},
			Rating{
				Name:   "4-Star",
				Rating: 50.0,
			},
		}
		got := calculateRatings(estab)
		if expected, actual := 2, len(got); expected != actual {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
		if expected, actual := want, got; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("fuzz", func(t *testing.T) {
		fn := func(a, b, c, d, e string) bool {
			ratings := [5]string{"a", "b", "c", "d", "e"}

			amount := rand.Intn(10000) + 100
			estab := make([]service.Establishment, amount)
			for i := 0; i < amount; i++ {
				estab[i] = service.Establishment{
					Name:   fmt.Sprintf("estab-%d", i),
					Rating: ratings[i%len(ratings)],
				}
			}
			got := calculateRatings(estab)
			want := make([]Rating, len(ratings))
			rating := func(index int) float64 {
				var (
					overflow   = amount % len(ratings)
					overflowf  = float64(overflow)
					numRatings = float64(len(ratings))
					amountf    = float64(amount)
					offset     = amountf - overflowf
				)
				if index < overflow {
					offset = amountf + (numRatings - overflowf)
				}
				return ((offset / numRatings) / amountf) * 100
			}
			for k := range ratings {
				want[k] = Rating{
					Name:   strings.Title(ratings[k]),
					Rating: rating(k),
				}
			}

			if expected, actual := want, got; !reflect.DeepEqual(expected, actual) {
				t.Errorf("expected: %v, actual: %v", expected, actual)
			}

			return true
		}

		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})
}

func TestRatingName(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		got := ratingName("")
		if expected, actual := "", got; expected != actual {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})

	t.Run("numeric", func(t *testing.T) {
		got := ratingName("341")
		if expected, actual := "341", got; expected != actual {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})

	t.Run("lowercase alpha", func(t *testing.T) {
		got := ratingName("alpha")
		if expected, actual := "Alpha", got; expected != actual {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})

	t.Run("uppercase alpha", func(t *testing.T) {
		got := ratingName("ALPHA")
		if expected, actual := "ALPHA", got; expected != actual {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})

	t.Run("1 to 5", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			got := ratingName(fmt.Sprintf("%d", i+1))
			want := fmt.Sprintf("%d-Star", i+1)
			if expected, actual := want, got; expected != actual {
				t.Errorf("expected: %s, actual: %s", expected, actual)
			}
		}
	})
}
