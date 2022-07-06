package service

import (
	entity "studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/division"
	posisirepo "studybuddy-backend-fast/api/repository/posisi"
)

type DivisionService interface {
	CreateDivision(divisi entity.Division) int
	FindAllDivisi() []entity.DivisiList
	Description(divid int, posid int) entity.DivisiAndPosisi
}

type divisionService struct {
	repo repository.DivisiRepo
	posisi posisirepo.PosisiRepo
}

func NewDivisionService(repo repository.DivisiRepo,posisi posisirepo.PosisiRepo) DivisionService {
	return &divisionService{repo,posisi}
}
func (s *divisionService) CreateDivision(divisi entity.Division) int  {
	return s.repo.CreateDivisi(divisi)
}

func (s *divisionService) FindAllDivisi() []entity.DivisiList {
	var data []entity.DivisiList
	var div entity.DivisiList
	divisi := s.repo.FindAllDivisi()
	for i := range divisi{
		posisi := s.posisi.FindAllPosisi(divisi[i].ID)
		div.DivisionID = divisi[i].ID
		div.Name = divisi[i].Name
		div.Position = posisi
		data = append(data,div)
	}
	return data
}
func (s *divisionService) Description(divid int, posid int) entity.DivisiAndPosisi{
	var data entity.DivisiAndPosisi
	divisi := s.repo.FindDivisiByID(divid)
	posisi := s.posisi.FindPosisi(divid,posid)
	data.DivisionName = divisi.Name
	data.PosisiName = posisi.Name
	data.RentangGaji = posisi.RentangGaji
	data.JobDesc = posisi.JobDesc
	data.Qualifikasi = posisi.Qualifikasi
	return data
}