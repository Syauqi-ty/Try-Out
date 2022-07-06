package entity

import (
	"time"
)

type Division struct{
	ID int `json:"id" gorm:"primary_key;auto_increment"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
type DivisiList struct {
	DivisionID int `json:"id"`
	Name string `json:"name"`
	Position []PosisiAll `json:"position"`
}
type DivisionAll struct {
	ID int `json:"id"`
	Name string `json:"name"`
}
type DivisiAndPosisi struct{
	DivisionName string `json:"divisi"`
	PosisiName string `json:"posisi"`
	RentangGaji string `json:"rentang_gaji"`
	JobDesc string `json:"job_desc"`
	Qualifikasi string `json:"qualifikasi"`
}