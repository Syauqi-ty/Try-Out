package entity

import (
	"time"
)

type Recruitment struct {
	ID int `json:"id" gorm:"primary_key;auto_increment"`
	ApplicantID int `json:"applicant_id"`
	Pendaftar Applicant `json:"pendaftar" gorm:"foreignKey:ApplicantID;references:ID"`
	Judul string `json:"judul"`
	Soal string `json:"soal"`
	Pilihan string `json:"pilihan"`
	Solusi string `json:"solusi"`
	CreatedAt time.Time `json:"created_at"`
}
type Soal struct {
	Judul string `json:"judul"`
	Soal string `json:"soal"`
	Pilihan string `json:"pilihan"`
	Solusi string `json:"solusi"`
}