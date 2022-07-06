package service

import (
	entity "studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/applicant"
	divisirepo "studybuddy-backend-fast/api/repository/division"
	posisirepo "studybuddy-backend-fast/api/repository/posisi"
)

type ApplicantService interface {
	FindAllApplicant() []entity.Applicant
	FindApplicantByID(id int) entity.Applicant
	CreateApplicant(applicant entity.Applicant) entity.Applicant
	CreateSolution(solution entity.Solution) entity.Solution
	UpdateApplicant(applicant entity.Applicant) entity.Applicant
	DeleteApplicant(applicant entity.Applicant)
	ListApplicant() []entity.ApplicantList
	Update(ID int,applicant entity.Applicant)
	Email(applicant entity.Applicant)
	ListApplicantNew() []entity.ApplicantListWithDiv
}

type applicantService struct {
	applicantrepo repository.ApplicantRepo
	divrepo divisirepo.DivisiRepo
	posrepo posisirepo.PosisiRepo
}

func NewApplicantService(rep repository.ApplicantRepo,divrep divisirepo.DivisiRepo,posrep posisirepo.PosisiRepo) ApplicantService {
	return &applicantService{rep,divrep,posrep}
}

func (s *applicantService) FindAllApplicant() []entity.Applicant {
	return s.applicantrepo.FindAllApplicant()
}

func (s *applicantService) FindApplicantByID(id int) entity.Applicant {
	return s.applicantrepo.FindApplicantByID(id)
}

func (s *applicantService) CreateApplicant(applicant entity.Applicant) entity.Applicant {
	return s.applicantrepo.CreateApplicant(applicant)
}

func (s *applicantService) UpdateApplicant(applicant entity.Applicant) entity.Applicant {
	s.applicantrepo.UpdateApplicant(applicant)
	return applicant
}

func (s *applicantService) DeleteApplicant(applicant entity.Applicant) {
	s.applicantrepo.DeleteApplicant(applicant)
}

func (s *applicantService) CreateSolution(solution entity.Solution) entity.Solution {
	return s.applicantrepo.CreateSolution(solution)
}
func (s *applicantService) ListApplicant() []entity.ApplicantList {
	return s.applicantrepo.ListApplicant()
}
func (s *applicantService) Update(ID int,applicant entity.Applicant) {
	s.applicantrepo.Update(ID,applicant)
}
func (s *applicantService) Email(applicant entity.Applicant) {
	s.applicantrepo.Email(applicant)
} 
func (s *applicantService)ListApplicantNew() []entity.ApplicantListWithDiv  {
	var wibu entity.ApplicantListWithDiv
	var ganteng []entity.ApplicantListWithDiv
	data := s.ListApplicant()
	for i := range data {
		divisi := s.divrepo.FindDivisiByID(data[i].DivisionID)
		posisi := s.posrepo.FindPosisi(data[i].DivisionID,data[i].PositionID)
		wibu.Name = data[i].Name
		wibu.Accepted = data[i].Accepted
		wibu.Email = data[i].Email
		wibu.Divisiname = divisi.Name
		wibu.Posisiname = posisi.Name
		wibu.Motivasi = data[i].Motivasi
		wibu.Phone = data[i].Phone
		wibu.ID = data[i].ID
		ganteng = append(ganteng,wibu)
	}
	return ganteng
}