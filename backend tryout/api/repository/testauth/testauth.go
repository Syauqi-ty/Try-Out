package repository

import (
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"

	"gorm.io/gorm"
)

type TestAuthRepo interface {
	CreateTestAuth(TestAuth entity.TestAuth) entity.TestAuth
	FindTestAuthByCreds(UsernameTO string, PasswordTO string) entity.TestAuth
	FindTestAuthByStudentIDAndTestID(StudentID int, TestID int) entity.TestAuth
	UpdateTestAuth(TestAuth map[string]interface{})
	LoginSpace(login entity.LoginSpace) entity.TestAuth
}

type database struct {
	conn *gorm.DB
}

func NewTestAuthRepo() TestAuthRepo {
	db := connection.Create()
	db.AutoMigrate(&entity.TestAuth{})
	return &database{db}
}

func (d *database) FindTestAuthByCreds(UsernameTO string, PasswordTO string) entity.TestAuth {
	var auth entity.TestAuth
	d.conn.Where("username_to = ? AND password_to = ?", UsernameTO, PasswordTO).First(&auth)
	return auth
}

func (d *database) FindTestAuthByStudentIDAndTestID(StudentID int, TestID int) entity.TestAuth {
	var auth entity.TestAuth
	d.conn.Where("student_id = ? AND test_id = ?", StudentID, TestID).First(&auth)
	return auth
}

func (d *database) CreateTestAuth(TestAuth entity.TestAuth) entity.TestAuth {
	d.conn.Create(&TestAuth)
	return TestAuth
}

func (d *database) UpdateTestAuth(TestAuth map[string]interface{}) {
	d.conn.Model(&entity.TestAuth{}).Where("test_id = ? AND student_id = ?", TestAuth["test_id"], TestAuth["student_id"]).Updates(TestAuth)
}

func (d *database) LoginSpace(login entity.LoginSpace) entity.TestAuth{
	var hehe entity.TestAuth
	result := d.conn.Table("test_auths").Where("username_to=? AND password_to=?",login.Username,login.Password).FirstOrInit(&hehe)
	if result.Error != nil{
		return entity.TestAuth{}
	} else{
		return hehe
	}
}
