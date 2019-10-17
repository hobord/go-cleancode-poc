package models

// PaginationOptions with limit and offset
type PaginationOptions struct {
	limit  int
	offset int
}

// OrderOption with columname and desc or asc bool
type OrderOption struct {
	column string
	desc   bool
}

// ListOptions with pagination and order
type ListOptions struct {
	pagination PaginationOptions
	order      OrderOption
}
