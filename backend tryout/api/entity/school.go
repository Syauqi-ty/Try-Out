package entity

import (
	"time"
)

type School struct {
	ID        uint64     `json:"id"`
	Slug      string     `json:"slug"`
	Name      string     `json:"name"`
	Logo      string     `json:"logo"`
	Image     string     `json:"image"`
	Operators []Operator `json:"operators" gorm:"many2many:school_operator"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updatedAt"`
}
