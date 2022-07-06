package entity

import "time"

type Circle struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Slug        string    `json:"slug"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Member1     string    `json:"member_1"`
	Member2     string    `json:"member_2"`
	Member3     string    `json:"member_3"`
	Member4     string    `json:"member_4"`
	Member5     string    `json:"member_5"`
	Member6     string    `json:"member_6"`
	Member7     string    `json:"member_7"`
	Member8     string    `json:"member_8"`
	Member9     string    `json:"member_9"`
	Member10    string    `json:"member_10"`
	Active      int       `json:"active" `
	CreatedAt   time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
