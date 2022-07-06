package service

import (
	entity "studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/posisi"
)

type PosisiService interface {
	CreatePosisi(posisi entity.Posisi) int
}

type posisiService struct {
	repo repository.PosisiRepo
}

func NewPosisiService(repo repository.PosisiRepo) PosisiService {
	return &posisiService{repo}
}
func (s *posisiService) CreatePosisi(posisi entity.Posisi) int  {
	return s.repo.CreatePosisi(posisi)
}