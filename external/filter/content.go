package filter

import (
	"github.com/tokopedia/go-filter-parser"
)

type GetFilter struct {
	Order     filter.String
	Direction filter.String
	Limit     filter.String
	Offset    filter.String
	Category  filter.String
}

func (f *GetFilter) FilterMap() filter.FilterMap {
	return filter.FilterMap{
		&f.Order:     "order",
		&f.Direction: "direction",
		&f.Limit:     "limit",
		&f.Offset:    "offset",
		&f.Category:  "category",
	}
}
