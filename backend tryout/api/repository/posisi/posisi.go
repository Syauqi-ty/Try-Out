package repository

import (
	conn "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
	"time"

	"gorm.io/gorm"
)

type PosisiRepo interface{
	CreatePosisi(posisi entity.Posisi) int
	FindAllPosisi(id int) []entity.PosisiAll
	FindPosisi(divid int,id int) entity.Posisi 
}
type posisiRepo struct {
	conn *gorm.DB
}

func NewPosisiRepo() PosisiRepo {
	db := conn.Create()
	db.AutoMigrate(entity.Posisi{})
	return &posisiRepo{db}
}

func (r *posisiRepo) CreatePosisi( posisi entity.Posisi) int {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	err := r.conn.Create(&posisi)
	r.conn.Model(&posisi).Update("created_at",now)
	if err != nil {
		return 0
	} else {
		return 1
	}
}
func (r *posisiRepo) FindAllPosisi(id int) []entity.PosisiAll {
	var posisi []entity.PosisiAll
	r.conn.Table("posisis").Where("division_id=?",id).Find(&posisi)
	return posisi
}
func (r *posisiRepo)FindPosisi(divid int,id int) entity.Posisi {
	var posisi entity.Posisi
	r.conn.Table("posisis").Where("id=? AND division_id=?",id,divid).Find(&posisi)
	return posisi
}

