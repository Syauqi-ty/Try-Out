package repository

import (
	"strings"
	"villa-akmali/api/connection"
	"villa-akmali/api/model"

	"gorm.io/gorm"
)



type CategoryRepo interface {
	CreateCategory(category model.Category) model.Category
	FindAll() []model.Category
	FindByID(id int) model.Category
}

type database struct {
	connection *gorm.DB
}

func NewCategoryRepo() CategoryRepo {
	db := connection.Create()
	db.AutoMigrate(&model.Category{})
	return &database{connection: db}
}

func (db *database) CreateCategory(category model.Category) model.Category  {
	if len(strings.TrimSpace(category.Name)) == 0 {
	}else{
		db.connection.Create(&category)
	}
	return category
}

func (db *database) FindAll() []model.Category  {
	var category []model.Category
	db.connection.Table("categories").Find(&category)
	return category
}

func (db *database) FindByID(id int) model.Category {
	var category model.Category
	if id == 0{
	} else{
		db.connection.Table("categories").Where("id = ?",id).First(&category)
	}
	return category
}