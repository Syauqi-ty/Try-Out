package entity

import (
	"time"
)

type Posisi struct {
	ID int `json:"id" gorm:"primary_key;auto_increment"`
	DivisionID int `json:"division_id" gorm:"not null"`
	Divisi Division `json:"divisi" gorm:"foreignKey:DivisionID;references:ID"`
	Name string `json:"name"`
	RentangGaji string `json:"rentang_gaji"`
	JobDesc string `json:"job_desc"`
	Qualifikasi string `json:"qualifikasi"`
	CreatedAt time.Time `json:"createdAt"`
}
type PosisiAll struct {
	ID int `json:"id"`
	Name string `json:"name"`
}