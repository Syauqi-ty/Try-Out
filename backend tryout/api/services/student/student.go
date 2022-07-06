package service

import (
	entity "studybuddy-backend-fast/api/entity"
	shalala "studybuddy-backend-fast/api/helper"
	repository "studybuddy-backend-fast/api/repository/student"
)

type StudentService interface {
	FindAllStudent(qstring map[string]interface{}, pagination map[string]int) []entity.Student
	FindByID(id int) entity.Student
	CreateStudent(newUser entity.Student) entity.Student
	UpdateStudent(newUserData entity.Student) entity.Student
	AuthStudent(username string, password string) (bool, entity.Student)
	DeleteStudent(student entity.Student)
	Forget(student entity.Student) uint64
}

type studentService struct {
	studentrepo repository.StudentRepo
}

func NewStudentService(repo repository.StudentRepo) StudentService {
	return &studentService{
		studentrepo: repo,
	}
}

func (service *studentService) FindAllStudent(qstring map[string]interface{}, pagination map[string]int) []entity.Student {
	return service.studentrepo.FindWithFilter(qstring, pagination)
}

func (service *studentService) FindByID(id int) entity.Student {
	return service.studentrepo.FindOneById(id)
}

func (service *studentService) CreateStudent(newUser entity.Student) entity.Student {
	hash := shalala.Encrypt(newUser.Password)
	newUser.Password = hash
	return service.studentrepo.CreateStudent(newUser)
}

func (service *studentService) UpdateStudent(newUserData entity.Student) entity.Student {
	hash := shalala.Encrypt(newUserData.Password)
	if newUserData.Password != "" {
		newUserData.Password = hash
	} else {
		newUserData.Password = ""
	}
	service.studentrepo.UpdateStudent(newUserData)
	return newUserData
}

func (service *studentService) DeleteStudent(student entity.Student) {
	service.studentrepo.DeleteStudent(student)
}

func (service *studentService) AuthStudent(username string, password string) (bool, entity.Student) {
	user := service.studentrepo.FindOneByUsernameOrEmail(username)
	auth := shalala.Verify(password, user.Password)
	return auth, user
}

func (service *studentService) Forget(student entity.Student) uint64 {
	return service.studentrepo.Forget(student)
}
