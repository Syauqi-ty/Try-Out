package repository

import (
	conn "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
	"time"

	"gorm.io/gorm"
)

type DivisiRepo interface{
	CreateDivisi(divisi entity.Division) int
	FindAllDivisi() []entity.DivisionAll
	FindDivisiByID(id int) entity.DivisionAll
}
type divisionRepo struct {
	conn *gorm.DB
}

func NewDivisiRepo() DivisiRepo {
	db := conn.Create()
	db.AutoMigrate(entity.Division{})
	return &divisionRepo{db}
}

func (r *divisionRepo) CreateDivisi( divisi entity.Division) int {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	r.conn.Create(&divisi)
	r.conn.Model(&divisi).Update("created_at",now)
	return 1
}
func (r *divisionRepo) FindAllDivisi() []entity.DivisionAll{
	var div []entity.DivisionAll
	r.conn.Table("divisions").Find(&div)
	return div
}
func (r *divisionRepo) FindDivisiByID(id int) entity.DivisionAll{
	var divisi entity.DivisionAll
	r.conn.Table("divisions").Where("id=?",id).Find(&divisi)
	return divisi
}

