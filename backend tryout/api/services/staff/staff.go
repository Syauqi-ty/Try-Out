package service

import (
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/staff"
)

type StaffService interface {
	FindAllStaff(qstring map[string]interface{}, pagination map[string]int) []entity.StaffMin
	FindById(id int) entity.Staff
	FindByAccess(Access string) []entity.Staff
	FindByAccessCheck(id string) string
	Save(entity.Staff) entity.Staff
	Update(staff entity.Staff)
	Delete(staff entity.Staff)
}

type staffService struct {
	staffrepo repository.StaffRepo
}

func NewStaffService(repo repository.StaffRepo) StaffService {
	return &staffService{
		staffrepo: repo,
	}
}

func (service *staffService) FindAllStaff(qstring map[string]interface{}, pagination map[string]int) []entity.StaffMin {
	return service.staffrepo.FindWithFilter(qstring, pagination)
}

func (service *staffService) FindById(id int) entity.Staff {
	return service.staffrepo.FindOneById(id)
}

func (service *staffService) FindByAccess(Access string) []entity.Staff {
	return service.staffrepo.FindByAccess(Access)
}

func (service *staffService) FindByAccessCheck(id string) string {
	var idkirim int
	idkirim, _ = strconv.Atoi(id)
	return service.staffrepo.FindByAccessCheck(idkirim)
}

func (service *staffService) Save(staff entity.Staff) entity.Staff {
	service.staffrepo.Save(staff)
	return staff
}

func (service *staffService) Update(staff entity.Staff) {
	service.staffrepo.Update(staff)
}

func (service *staffService) Delete(staff entity.Staff) {
	service.staffrepo.Delete(staff)
}
