package service

import (
	entity "studybuddy-backend-fast/api/entity"
	shalala "studybuddy-backend-fast/api/helper"
	repository "studybuddy-backend-fast/api/repository/parent"
)

type ParentService interface {
	FindAllParent(qstring map[string]interface{}, pagination map[string]int) []entity.Parent
	FindParentByID(id int) entity.Parent
	CreateParent(parent entity.Parent) entity.Parent
	UpdateParent(parent entity.Parent)
	DeleteParent(parent entity.Parent)
	Login(parent entity.Parent) entity.ParentLogin
}

type parentService struct {
	repo repository.ParentRepo
}

func NewParentService(repo repository.ParentRepo) ParentService {
	return &parentService{repo}
}

func (s *parentService) FindAllParent(qstring map[string]interface{}, pagination map[string]int) []entity.Parent {
	return s.repo.FindParentWithFilter(qstring, pagination)
}

func (s *parentService) FindParentByID(id int) entity.Parent {
	return s.repo.FindParentByID(id)
}
func (s *parentService) Login(parent entity.Parent) entity.ParentLogin{
	var parentlogin entity.ParentLogin
	var you []uint64
	id := s.repo.Login(parent)
	parents := s.repo.FindParentByID(int(id))
	parentlogin.ID = parents.ID
	for t := range parent.Students{
		you := append(you,parent.Students[t].ID)
		parentlogin.StudentID = you
	}
	return parentlogin
}

func (s *parentService) CreateParent(parent entity.Parent) entity.Parent {
	parent.Password = shalala.Encrypt(parent.Password)
	return s.repo.CreateParent(parent)
}

func (s *parentService) UpdateParent(parent entity.Parent) {
	if parent.Password != "" {
		parent.Password = shalala.Encrypt(parent.Password)
	}
	s.repo.UpdateParent(parent)
}

func (s *parentService) DeleteParent(parent entity.Parent) {
	s.repo.DeleteParent(parent)
}
