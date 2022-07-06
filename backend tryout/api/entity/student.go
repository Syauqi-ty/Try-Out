package entity

import "time"

type Student struct {
	ID           uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Username     string    `json:"username" gorm:"unique"`
	Email        string    `json:"email" gorm:"unique"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	StudentClass string    `json:"student_class" gorm"column:student_class"`
	Password     string    `json:"password"`
	Brights      int       `json:"brights"`
	Target       int       `json:"target"`
	School       string    `json:"school"`
	CircleID     int       `json:"cirle_id" gorm:"column:circle_id"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
