package service

import (
	entity "studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/question"
)

type QuestionService interface {
	FindAllQuestion() []entity.Question
	FindQuestionByID(id int) entity.Question
	FindTestQuestion(slug string) []entity.Question
	CreateQuestion(question entity.Question) entity.Question
	UpdateQuestion(question entity.Question) int
	DeleteQuestion(id int)
	TemporaryDelete(id int)

	FindAllQuestionMin() []entity.QuestionMin
}

type questionService struct {
	repo repository.QuestionRepo
}

func NewQuestionService(questionRepo repository.QuestionRepo) QuestionService {
	return &questionService{repo: questionRepo}
}

func (s *questionService) FindAllQuestion() []entity.Question {
	return s.repo.FindAllQuestion()
}

func (s *questionService) FindAllQuestionMin() []entity.QuestionMin {
	return s.repo.FindAllQuestionMin()
}

func (s *questionService) FindQuestionByID(id int) entity.Question {
	return s.repo.FindQuestionByID(id)
}

func (s *questionService) FindTestQuestion(slug string) []entity.Question {
	print(slug)
	return s.repo.FindAllQuestion()
}

func (s *questionService) CreateQuestion(question entity.Question) entity.Question {
	return s.repo.CreateQuestion(question)
}

func (s *questionService) UpdateQuestion(question entity.Question) int {
	return s.repo.UpdateQuestion(question)
}

func (s *questionService) DeleteQuestion(id int) {
	s.repo.DeleteQuestion(id)
}
func (s *questionService) TemporaryDelete(id int) {
	s.repo.TemporaryDelete(id)
}
