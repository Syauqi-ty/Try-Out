package service

import (
	entity "studybuddy-backend-fast/api/entity"
	shalala "studybuddy-backend-fast/api/helper"
	operatorRepository "studybuddy-backend-fast/api/repository/operator"
	parentRepository "studybuddy-backend-fast/api/repository/parent"
	staffRepository "studybuddy-backend-fast/api/repository/staff"
	studentRepository "studybuddy-backend-fast/api/repository/student"
	validators "studybuddy-backend-fast/api/validators"
)

var (
	studentRepo  studentRepository.StudentRepo   = studentRepository.NewStudentRepo()
	staffRepo    staffRepository.StaffRepo       = staffRepository.NewStaffRepo()
	parentRepo   parentRepository.ParentRepo     = parentRepository.NewParentRepo()
	operatorRepo operatorRepository.OperatorRepo = operatorRepository.NewOperatorRepo()
)

type User struct {
	Staff   entity.Staff
	Student entity.Student
	Parent  entity.Parent
}

func Auth(creds validators.AuthBody) (bool, User) {
	var user User

	user.Staff = staffRepo.FindOneByUsernameOrEmail(creds.Username)
	if user.Staff.ID != 0 {
		return shalala.Verify(creds.Password, user.Staff.Password), user
	}

	user.Student = studentRepo.FindOneByUsernameOrEmail(creds.Username)
	if user.Student.ID != 0 {
		return shalala.Verify(creds.Password, user.Student.Password), user
	}

	user.Parent = parentRepo.FindParentByUsername(creds.Username)
	if user.Parent.ID != 0 {
		return shalala.Verify(creds.Password, user.Parent.Password), user
	}

	return false, user
}

func OperatorAuth(creds validators.AuthBody) (bool, entity.Operator) {
	if operator := operatorRepo.FindOperatorByUsername(creds.Username); operator.ID != 0 {
		return shalala.Verify(creds.Password, operator.Password), operator
	} else {
		return false, operator
	}
}
