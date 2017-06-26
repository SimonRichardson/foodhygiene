package query

import (
	"net/url"

	"github.com/pkg/errors"
)

// EstablishmentsQueryParams defines all the dimensions of a query.
type EstablishmentsQueryParams struct {
	LocalID string
}

// DecodeFrom populates a EstablishmentsQueryParams from a URL.
func (p *EstablishmentsQueryParams) DecodeFrom(u *url.URL, rb queryBehavior) error {
	// Required depending on the query behavior
	p.LocalID = u.Query().Get("local_id")
	if p.LocalID == "" && rb == queryRequired {
		return errors.New("error reading/parsing 'local_id' (required) query")
	}
	return nil
}
