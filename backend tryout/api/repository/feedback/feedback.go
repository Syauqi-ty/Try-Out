package repository

import (
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"

	"gorm.io/gorm"
)

type konak struct {
	connection *gorm.DB
}
type FeedbackRepo interface {
	FindAllFeedback() []entity.Feedback
	CreateFeedBack(f entity.Feedback) entity.Feedback
}

func NewFeedbackRepo() FeedbackRepo {
	// koneksi ke db gais
	db := connection.Create()
	db.AutoMigrate(&entity.Feedback{})

	// end contoh
	return &konak{
		connection: db,
	}
}

func (db *konak) FindAllFeedback() []entity.Feedback {
	var feedback []entity.Feedback
	db.connection.Find(&feedback)
	return feedback
}

func (db *konak) CreateFeedBack(f entity.Feedback) entity.Feedback{
	db.connection.Create(&f)
	return f
}