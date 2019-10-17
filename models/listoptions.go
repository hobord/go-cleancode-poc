package models

import (
	"fmt"
)

// PaginationOptions with limit and offset
type PaginationOptions struct {
	Limit  int
	Offset int
}

// OrderOption with columname and desc or asc bool
type OrderOption struct {
	Column string
	Desc   bool
}

// ListOptions with pagination and order
type ListOptions struct {
	Pagination PaginationOptions
	Order      []OrderOption
}

// MysqlQueryStr make a mysql query str from options
func (listOptions *ListOptions) MysqlQueryStr() string {
	query := ""
	if len(listOptions.Order) > 0 {
		length := len(listOptions.Order)
		query = query + " ORDER BY "
		orders := ""
		for i, column := range listOptions.Order {
			order := "ASC"

			if column.Desc {
				order = "DESC"
			}

			coma := ","
			if i == length-1 {
				coma = ""
			}
			orders = orders + column.Column + " " + order + coma
		}
		query = query + orders
	}

	if listOptions.Pagination.Limit > 0 {
		offset := ""
		if listOptions.Pagination.Offset > 0 {
			offset = "," + fmt.Sprintf("%d", listOptions.Pagination.Offset)
		}
		limit := fmt.Sprintf("%d", listOptions.Pagination.Limit)
		query = query + " LIMIT " + limit + offset
	}

	return query
}
