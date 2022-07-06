package repository

import (
	"villa-akmali/api/connection"
	"villa-akmali/api/model"

	"gorm.io/gorm"
)


type YearRepo interface {
	CreateYear(year model.Year)
	FindAll() []model.Year
}

type database struct {
	connection *gorm.DB
}

func NewYearRepo() YearRepo {
	db := connection.Create()
	db.AutoMigrate(&model.Year{})
	return &database{connection: db}
}

func (db *database) CreateYear(year model.Year) {
	data := db.connection.Table("years").Where("year = ?",year.Year).Find(&year)
	if data.RowsAffected == 0{
		db.connection.Create(&year)
	}
}

func (db *database) FindAll() []model.Year {
	var year []model.Year
	db.connection.Table("years").Find(&year)
	return year
}