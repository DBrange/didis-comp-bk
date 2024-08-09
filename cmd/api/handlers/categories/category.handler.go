package handlers

import ports "github.com/DBrange/didis-comp-bk/domains/category/ports/drivers"

type Handler struct {
	category ports.ForCategory
}

func NewHandlerCategory(category ports.ForCategory) *Handler {
	return &Handler{
		category: category,
	}
}
