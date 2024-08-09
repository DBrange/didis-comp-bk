package services

import ports "github.com/DBrange/didis-comp-bk/domains/category/ports/drivens"

type CategoryService struct {
	categoryQueryer ports.ForQueryingCategory
}

func NewCategoryService(categoryQueryer ports.ForQueryingCategory) *CategoryService {
	return &CategoryService{
		categoryQueryer: categoryQueryer,
	}
}
