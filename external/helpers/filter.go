package helpers

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/tokopedia/go-filter-parser"
	f "go-hdflex/external/filter"
	"log/slog"
	"strconv"
)

func ApplyParams(q *goqu.SelectDataset, filter *f.GetFilter) *goqu.SelectDataset {
	q = applyOrderParam(q, filter)
	q = applyLimitParam(q, filter.Limit)
	q = applyOffsetParam(q, filter.Offset)
	q = applyCategoryParam(q, filter.Category)

	return q
}

func applyOrderParam(q *goqu.SelectDataset, filter *f.GetFilter) *goqu.SelectDataset {
	field := "id"

	if filter.Order.Value != "" {
		field = filter.Order.Value
	}

	if filter.Direction.Value == "asc" {
		q = q.Order(goqu.I(field).Asc())
	} else if filter.Direction.Value == "desc" {
		q = q.Order(goqu.I(field).Desc())
	} else {
		q = q.Order(goqu.I(field).Asc())
	}

	return q
}

func applyLimitParam(q *goqu.SelectDataset, filter filter.String) *goqu.SelectDataset {
	var err error
	limit := 20

	if filter.Value != "" {
		limit, err = strconv.Atoi(filter.Value)
		if err != nil {
			slog.Error("Error parsing limit", err)
			limit = 20
		}
	}

	q = q.Limit(uint(limit))

	return q
}

func applyCategoryParam(q *goqu.SelectDataset, filter filter.String) *goqu.SelectDataset {
	if filter.Value != "" && filter.Value != "0" {
		q = q.Where(goqu.Ex{"contents.content_type": filter.Value})
	}

	return q
}

func applyOffsetParam(q *goqu.SelectDataset, filter filter.String) *goqu.SelectDataset {
	var err error
	offset := 0

	if filter.Value != "" {
		offset, err = strconv.Atoi(filter.Value)
		if err != nil {
			slog.Error("Error parsing offset", err)
			offset = 0
		}
	}

	q = q.Offset(uint(offset))

	return q
}
