package model

import "time"

type Order struct {
	ID         int      `json:"id" gorm:"primary_key;auto_increment"`
	Judul string `json:"judul"`
	Keterangan string   `json:"keterangan"`
	CategoryID int      `json:"category_id"`
	Kategori   Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	UserID int `json:"user_id"`
	User   User `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Type       string   `json:"type"`
	Image string `json:"image"`
	Jumlah     float64   `json:"jumlah" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"created_at"`
}

type Orderlist struct {
	ID         int      `json:"id"`
	Judul string `json:"judul"`
	Keterangan string   `json:"keterangan"`
	CategoryID int      `json:"category_id"`
	UserID     int `json:"user_id"`
	Type       string   `json:"type"`
	Image string `json:"image"`
	Jumlah     float64   `json:"jumlah" gorm:"type:text"`
	CreatedAt  time.Time `json:"created_at"`
}
type Orderlist2 struct {
	ID         int      `json:"id"`
	Keterangan string   `json:"keterangan"`
	CategoryID int      `json:"category_id"`
	UserID int `json:"user_id"`
	Type       string   `json:"type"`
	Image string `json:"image"`
	Jumlah     float64   `json:"jumlah" gorm:"type:text"`
	CreatedAt  time.Time `json:"created_at"`
}

type Query struct {
	Waktu string `json:"name"`
	Input float64 `json:"Income"`
	Output float64 `json:"Expense"`
}
type QuerySum struct {
	Input float64 `json:"input"`
	Output float64 `json:"output"`
}
type Graph struct {
	Jumlah float64 `json:"jumlah"`
	Waktu string `json:"waktu"`
}

type Graph2 struct {
	Jumlah float64 `json:"jumlah"`
	Waktu string `json:"waktu"`
}
type AllOrder struct {
	ID int `json:"id"`
	Judul string `json:"judul"`
	CategoryName string      `json:"category"`
	Username   string `json:"user"`
	Jumlah float64 `json:"jumlah"`
}

type ById struct {
	ID int `json:"id"`
	Judul string `json:"judul"`
	Keterangan string `json:"keterangan"`
	CategoryName string      `json:"category"`
	Username   string `json:"user"`
	Type string `json:"type"`
	Image string `json:"image"`
	Jumlah float64 `json:"jumlah"`
	CreatedAt time.Time  `json:"created_at"`	
}

type Pagination struct {
	Limit int `json:"limit"`
	Page int `json:"page"`
	Sort string `json:"sort"`
}

type QueryFilter struct {
	StartedAt string `json:"started_at"`
	EndsAt string `json:"ends_at"`
	Category string `json:"category"`
}