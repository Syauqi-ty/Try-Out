package service

import (
	entity "studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/school"
)

type SchoolService interface {
	FindSchoolBySlug(slug string) entity.School
	FindSchoolByID(id int) entity.School
	FindAllSchool() []entity.School
	CreateSchool(school entity.School) entity.School
	UpdateSchool(school entity.School) entity.School
	DeleteSchool(school entity.School)
}

type schoolService struct {
	repo repository.SchoolRepo
}

func NewSchoolService(repo repository.SchoolRepo) SchoolService {
	return &schoolService{repo}
}

func (s *schoolService) FindSchoolBySlug(slug string) entity.School {
	return s.repo.FindSchoolBySlug(slug)
}

func (s *schoolService) FindSchoolByID(id int) entity.School {
	return s.repo.FindSchoolByID(id)
}

func (s *schoolService) FindAllSchool() []entity.School {
	return s.repo.FindAllSchool()
}

func (s *schoolService) CreateSchool(school entity.School) entity.School {
	s.repo.CreateSchool(school)
	return school
}

func (s *schoolService) UpdateSchool(school entity.School) entity.School {
	s.repo.UpdateSchool(school)
	return s.repo.FindSchoolByID(int(school.ID))
}

func (s *schoolService) DeleteSchool(school entity.School) {
	if school.ID != 0 {
		s.repo.DeleteSchool(school)
	}
}
