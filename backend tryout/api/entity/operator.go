package entity

import "time"

type Operator struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	School    string    `json:"school"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
