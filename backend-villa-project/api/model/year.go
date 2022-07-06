package model

type Year struct {
	ID   int `json:"id" gorm:"primary_key;auto_increment"`
	Year int `json:"year" gorm:"unique"`
}