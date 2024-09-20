package meta

import (
	"os"
	"strconv"
)

type Meta struct {
	TotalCount int `json:"total_count"`
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
}

func NewMeta(page, perPage, total int, pagLimDef string) (*Meta, error) {

	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(os.Getenv(pagLimDef))

		if err != nil {
			return nil, err
		}
	}

	pageCount := 0

	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}

	if page < 1 {
		page = 1
	}

	return &Meta{
		Page:       page,
		PerPage:    perPage,
		TotalCount: total,
	}, nil
}

func (o *Meta) Offset() int {
	return (o.Page - 1) * o.PerPage
}

func (p *Meta) Limit() int {
	return p.PerPage
}
