package entity

import (
	"time"
)

type Question struct {
	ID            uint64     `json:"id" gorm:"primary_key;auto_increment"`
	Name          string     `json:"name"`
	Type          string     `json:"type"`
	Question      string     `json:"question"`
	Answer        string     `json:"answer"`
	Duration      int        `json:"duration"`
	CreatorID     uint64     `json:"creator_id"`
	Creator       Staff   `json:"creator" gorm:"foreignKey:CreatorID;references:ID;"`
	UsedFor       []TestBare `json:"used_for" gorm:"many2many:test_question;foreignKey:ID;joinForeignKey:question_id;References:ID;JoinReferences:test_id"`
	CreatedAt     time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt     time.Time  `json:"deletedAt" gorm:"column:deletedAt"`
	UsedAt        time.Time  `json:"usedAt" gorm:"column:usedAt"`
}

type QuestionMin struct {
	ID            uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	Question      string    `json:"question"`
	Answer        string    `json:"answer"`
	Duration      int       `json:"duration"`
	CreatorID     uint64    `json:"creator_id"`
	LastUpdatorID uint64    `json:"last_updator_id"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt     time.Time `json:"deletedAt" gorm:"column:deletedAt"`
	UsedAt        time.Time `json:"usedAt" gorm:"column:usedAt"`
}

type QuestionAndSolution struct {
	ID            uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	Question      string    `json:"question"`
	Content       string    `json:"content"`
	Answer        string    `json:"answer"`
	Duration      int       `json:"duration"`
	CreatorID     uint64    `json:"creator_id"`
	LastUpdatorID uint64    `json:"last_updator_id"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt     time.Time `json:"deletedAt" gorm:"column:deletedAt"`
	UsedAt        time.Time `json:"usedAt" gorm:"column:usedAt"`
}
