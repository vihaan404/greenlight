package data

import (
	"strings"

	"github.com/vihaan404/greenlight/internal/validator"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

func (f Filters) limit() int {
	return f.PageSize
}

func (f Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}

func (f Filters) sortColumn() string {
	if f.Sort == "" {
		return "id"
	}
	for _, safeValue := range f.SortSafelist {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort paramenter: " + f.Sort)
}

func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than 0 ")
	v.Check(f.Page <= 10_000_000, "page", "must be  maximum of 10 million ")
	v.Check(f.Page > 0, "page_size", "must be greater than 0 ")
	v.Check(f.PageSize <= 100, "page_size", "must be  a maximum of 100 ")
	v.Check(validator.In(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}
