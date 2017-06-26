package query

import (
	"fmt"
	"sort"
	"strings"

	"github.com/SimonRichardson/foodhygiene/pkg/service"
)

// Rating is a type that defines a key value pair of a name and a rating value
// as a percentage or string for specific cases i.e. "Exempt"
type Rating struct {
	Name   string
	Rating float64
}

func calculateRatings(establishments []service.Establishment) []Rating {
	// So ratings is actually quite loose, you can have a lot of various values
	// for the key, which makes things a bit more complicated.
	var (
		total  int
		values = map[string]int{}
	)

	// Go through and increment all the values found by ratings.
	for _, v := range establishments {
		values[v.Rating]++
		total++
	}

	// Now convert them to percentages
	ratings := make([]Rating, 0, len(values))
	for k, v := range values {
		ratings = append(ratings, Rating{
			Name:   nameRating(k),
			Rating: float64(v) / float64(total),
		})
	}

	// Now let's make sure we sort them into some decent order
	sort.Slice(ratings, func(i, j int) bool { return ratings[i].Name < ratings[j].Name })

	return ratings
}

// nameRatings converts values into correctly expected rating values
// i.e. "3" == "3-star" and "pass" == "Pass"
func nameRating(name string) string {
	switch name {
	case "1", "2", "3", "4", "5":
		return fmt.Sprintf("%s-star", name)
	default:
		return strings.Title(name)
	}
}
