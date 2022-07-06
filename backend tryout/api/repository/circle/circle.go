package repository

import (
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var timezone string = viper.GetString("timezone")

type CircleRepo interface {
	FindAllCircle() []entity.Circle
	FindCircleByID(id int) entity.Circle
	CreateCircle(circle entity.Circle) entity.Circle
	UpdateCircle(circle entity.Circle)
	DeleteCircle(circle entity.Circle)
}

type database struct {
	connection *gorm.DB
}

func NewCircleRepo() CircleRepo {
	db := connection.Create()
	db.AutoMigrate(&entity.Circle{})
	return &database{connection: db}
}

func (db *database) CreateCircle(circle entity.Circle) entity.Circle {
	loc, _ := time.LoadLocation(timezone)
	circle.CreatedAt = time.Now().In(loc)
	circle.UpdatedAt = time.Now().In(loc)
	db.connection.Create(&circle)
	return circle
}

func (db *database) FindAllCircle() []entity.Circle {
	var circle []entity.Circle
	db.connection.Find(&circle)
	return circle
}

func (db *database) FindCircleByID(id int) entity.Circle {
	var circle entity.Circle
	db.connection.First(&circle, id)
	return circle
}

func (db *database) UpdateCircle(circle entity.Circle) {
	loc, _ := time.LoadLocation(timezone)
	circle.UpdatedAt = time.Now().In(loc)
	db.connection.Save(&circle)
}

func (db *database) DeleteCircle(circle entity.Circle) {
	db.connection.Delete(&circle)
}
