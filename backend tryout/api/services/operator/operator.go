package service

import (
	entity "studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/operator"
	shalala "studybuddy-backend-fast/api/helper"
)

type OperatorService interface {
	FindOperator(params map[string]interface{}, pagination map[string]int) []entity.Operator
	FindOperatorByID(id int) entity.Operator
	CreateOperator(operator entity.Operator) entity.Operator
	UpdateOperator(operator entity.Operator)
	DeleteOperator(operator entity.Operator)
}

type operatorService struct {
	repo repository.OperatorRepo
}

func NewOperatorRepo(repo repository.OperatorRepo) OperatorService {
	return &operatorService{repo}
}

func (s *operatorService) FindOperator(params map[string]interface{}, pagination map[string]int) []entity.Operator {
	return s.repo.FindOperatorWithFilter(params, pagination)
}

func (s *operatorService) FindOperatorByID(id int) entity.Operator {
	return s.repo.FindOperatorByID(id)
}

func (s *operatorService) CreateOperator(operator entity.Operator) entity.Operator {
	operator.Password = shalala.Encrypt(operator.Password)
	return s.repo.CreateOperator(operator)
}

func (s *operatorService) UpdateOperator(operator entity.Operator) {
	if operator.Password != "" {
		operator.Password = shalala.Encrypt(operator.Password)
	}
	s.repo.UpdateOperator(operator)
}

func (s *operatorService) DeleteOperator(operator entity.Operator) {
	s.repo.DeleteOperator(operator)
}
