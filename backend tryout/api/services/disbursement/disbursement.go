package service

import (
	"studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/disbursement"
)

type DisbursementService interface {
	FindAllDisbursement() []entity.Disbursement
	CreateDisbursement(disbursement entity.Disbursement) entity.Disbursement
}

type disbursementService struct {
	disbursementrepo repository.DisbursementRepo
}

func NewDisbursementService(repo repository.DisbursementRepo) DisbursementService {
	return &disbursementService{
		disbursementrepo: repo,
	}
}

func (s *disbursementService) FindAllDisbursement() []entity.Disbursement {
	return s.disbursementrepo.FindAllDisbursement()
}

func (s *disbursementService) CreateDisbursement(disbursement entity.Disbursement) entity.Disbursement {
	return s.disbursementrepo.CreateDisbursement(disbursement)
}
