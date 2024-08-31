package services

import ports "github.com/DBrange/didis-comp-bk/domains/category/ports/drivens"

type CategoryService struct {
	categoryQuerier ports.ForQueryingCategory
}

func NewCategoryService(categoryQuerier ports.ForQueryingCategory) *CategoryService {
	return &CategoryService{
		categoryQuerier: categoryQuerier,
	}
}
