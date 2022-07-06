package service

import (
	entity "studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/recruitment"
	apprepo "studybuddy-backend-fast/api/repository/applicant"
)

type RecruitmentService interface {
	CreateSoal(applicant entity.Applicant,soal []entity.Recruitment) int
	FindSoalByApplicantID(ID int) []entity.Soal
}

type recruitmentService struct {
	repo repository.RecruitmentRepo
	app apprepo.ApplicantRepo
}

func NewRecruitmentService(repo repository.RecruitmentRepo,app apprepo.ApplicantRepo) RecruitmentService {
	return &recruitmentService{repo,app}
}
func (s *recruitmentService) CreateSoal(applicant entity.Applicant,soal []entity.Recruitment) int {
	 app := s.app.CreateApplicant(applicant)
	 for i := range soal{
		 soal[i].ApplicantID = app.ID
	 }
	err :=  s.repo.CreateSoal(soal)
	 if err != nil {
		 return 0
	 } else {
		 return 1
	 }
}
func (s *recruitmentService) FindSoalByApplicantID(ID int) []entity.Soal {
	return s.repo.FindSoalByApplicantID(ID)
}
