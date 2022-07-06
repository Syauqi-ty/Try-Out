package service

import (
	"villa-akmali/api/model"
	repository "villa-akmali/api/repository/category"
)


type CategoryService interface {
	CreateCategory(category model.Category) model.Category 
	FindAll() []model.Category
}

type categoryService struct {
	categoryrepo repository.CategoryRepo
}

func NewCategoryService(repo repository.CategoryRepo) CategoryService {
	return &categoryService{
		categoryrepo: repo,
	}
}

func (service *categoryService) CreateCategory(category model.Category) model.Category {
	return service.categoryrepo.CreateCategory(category)
}

func (service *categoryService) FindAll() []model.Category {
	return service.categoryrepo.FindAll()
}