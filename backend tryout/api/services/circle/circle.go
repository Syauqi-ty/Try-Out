package service

import (
	entity "studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/circle"
)

type CircleService interface {
	FindAllCircle() []entity.Circle
	FindCircleByID(id int) entity.Circle
	CreateCircle(circle entity.Circle) entity.Circle
	UpdateCircle(circle entity.Circle) entity.Circle
	DeleteCircle(circle entity.Circle)
}

type circleService struct {
	circlerepo repository.CircleRepo
}

func NewCircleService(repo repository.CircleRepo) CircleService {
	return &circleService{
		circlerepo: repo,
	}
}

func (s *circleService) FindAllCircle() []entity.Circle {
	return s.circlerepo.FindAllCircle()
}

func (s *circleService) FindCircleByID(id int) entity.Circle {
	return s.circlerepo.FindCircleByID(id)
}

func (s *circleService) CreateCircle(circle entity.Circle) entity.Circle {
	return s.circlerepo.CreateCircle(circle)
}

func (s *circleService) UpdateCircle(circle entity.Circle) entity.Circle {
	s.circlerepo.UpdateCircle(circle)
	return circle
}

func (s *circleService) DeleteCircle(circle entity.Circle) {
	s.circlerepo.DeleteCircle(circle)
}
