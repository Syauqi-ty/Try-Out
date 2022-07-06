package repository

import (
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
	"time"

	"gorm.io/gorm"
)

type RecruitmentRepo interface{
	CreateSoal(soal []entity.Recruitment)[]entity.Recruitment
	FindSoalByApplicantID(ID int) []entity.Soal
}

type recruitmentRepo struct {
	conn *gorm.DB
}

func NewRecruitmentRepo() RecruitmentRepo {
	conn := connection.Create()
	conn.AutoMigrate(&entity.Recruitment{})
	return &recruitmentRepo{conn}
}

func (s *recruitmentRepo) CreateSoal(soal []entity.Recruitment)[]entity.Recruitment{
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	s.conn.Create(soal).Update("created_at",now)
	return soal
}

func (s *recruitmentRepo) FindSoalByApplicantID(ID int) []entity.Soal {
	var soal []entity.Soal
	s.conn.Table("recruitments").Where("applicant_id=?",ID).Find(&soal)
	return soal
}

