package service

import (
	entity "studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/pref"
)

type PrefService interface {
	FindAllUni() []entity.Uni
	FindAllFUni() []entity.FUni
	FindAllUniMin() []entity.UniMin
	FindUniByID(id int) entity.Uni
	FindUniMinByID(id int) entity.UniMin
	FindFUniByID(id int) entity.FUni

	FindProdiOfUni(ProdiID int, UniID int) entity.FUni
}

type prefService struct {
	repo repository.PrefRepo
}

func NewPrefService(repo repository.PrefRepo) PrefService {
	return &prefService{repo}
}

//////////////
// USECASES //
//////////////

func (s *prefService) FindAllFUni() []entity.FUni {
	var funis []entity.FUni

	for _, uni := range s.repo.FindAllUni() {
		studies := entity.SciSoc{Sci: uni.Sci, Soc: uni.Soc}

		funis = append(funis, entity.FUni{
			ID:      uni.ID,
			UniID:   uni.UniID,
			Name:    uni.Name,
			Studies: studies,
		})
	}

	return funis
}

func (s *prefService) FindFUniByID(id int) entity.FUni {
	uni := s.repo.FindUniByUniID(id)
	studies := entity.SciSoc{Sci: uni.Sci, Soc: uni.Soc}
	return entity.FUni{
		ID:      uni.ID,
		UniID:   uni.UniID,
		Name:    uni.Name,
		Studies: studies,
	}
}

func (s *prefService) FindProdiOfUni(ProdiID int, UniID int) entity.FUni {
	uni := s.repo.FindUniMinByUniID(UniID)
	prodi := s.repo.FindProdiOfUni(ProdiID, UniID)

	var studies entity.SciSoc
	if prodi.Type == "sci" {
		studies = entity.SciSoc{Sci: []entity.Prodi{prodi}}
	} else {
		studies = entity.SciSoc{Soc: []entity.Prodi{prodi}}
	}

	return entity.FUni{
		ID:      uni.ID,
		UniID:   uni.UniID,
		Name:    uni.Name,
		Studies: studies,
	}
}

func (s *prefService) FindAllUni() []entity.Uni {
	return s.repo.FindAllUni()
}

func (s *prefService) FindAllUniMin() []entity.UniMin {
	return s.repo.FindAllUniMin()
}

func (s *prefService) FindUniByID(id int) entity.Uni {
	return s.repo.FindUniByID(id)
}

func (s *prefService) FindUniMinByID(id int) entity.UniMin {
	return s.repo.FindUniMinByID(id)
}
