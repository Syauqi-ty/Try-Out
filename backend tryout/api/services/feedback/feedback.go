package service

import (
	entity "studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/feedback" 
)
type FeedbackService interface {
	FindAllFeedback() []entity.Feedback
	CreateFeedback(f entity.Feedback) entity.Feedback
}
type feedbackService struct {
	repo repository.FeedbackRepo
}

func NewFeedbackService(repo repository.FeedbackRepo) FeedbackService {
	return &feedbackService{repo}
}

func (s *feedbackService) FindAllFeedback() []entity.Feedback {
	return s.repo.FindAllFeedback()
}

func (s *feedbackService) CreateFeedback(f entity.Feedback) entity.Feedback{
	return s.repo.CreateFeedBack(f)
}