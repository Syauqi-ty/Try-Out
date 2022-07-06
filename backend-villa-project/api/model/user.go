package model

import "time"

type User struct {
	ID        int    `json:"id" gorm:"primary_key;auto_increment"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Phone string `json:"phone" gorm:"unique"`
	Password  string `json:"password"`
	Level     int    `json:"level"`
	CreatedAt time.Time `json:"created_at"`
}

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}