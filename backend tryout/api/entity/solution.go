package entity

import "time"

type Solution struct {
	ID         uint64    `json:"id" gorm:"foreign_key;auto_increment"`
	QuestionID uint64    `json:"question_id"`
	Pertanyaan Question  `json:"question" gorm:"foreignKey:QuestionID;references:ID;"`
	Content    string    `json:"content"`
	CreatorID  uint64    `json:"creator_id"`
	Creator    Staff     `json:"creator" gorm:"foreignKey:CreatorID;references:ID;"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}

type SolutionMin struct {
	ID         uint64    `json:"id" gorm:"foreign_key;auto_increment"`
	QuestionID uint64    `json:"question_id"`
	Content    string    `json:"content"`
	CreatorID  uint64    `json:"creator_id"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
