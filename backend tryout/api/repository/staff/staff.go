package repository

import (
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"

	"time"

	"gorm.io/gorm"
)

type StaffRepo interface {
	Save(staff entity.Staff)
	FindAll() []entity.Staff
	FindOneById(id int) entity.Staff
	FindOneByUsernameOrEmail(username string) entity.Staff
	FindByAccess(Access string) []entity.Staff
	FindByAccessCheck(id int) string
	Update(staff entity.Staff)
	Delete(staff entity.Staff)
	FindWithFilter(qstring map[string]interface{}, pagination map[string]int) []entity.StaffMin
}

type konak struct {
	connection *gorm.DB
}

func NewStaffRepo() StaffRepo {
	// koneksi ke db gais

	db := connection.Create()
	db.AutoMigrate(&entity.Staff{})

	// end contoh

	return &konak{
		connection: db,
	}
}

func (db *konak) Save(staff entity.Staff) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	db.connection.Create(&staff)
	db.connection.Model(&staff).Update("created_at", now)
}

func (db *konak) FindAll() []entity.Staff {
	var alluser []entity.Staff
	db.connection.Find(&alluser)
	return alluser
}

func (db *konak) FindWithFilter(qstring map[string]interface{}, pagination map[string]int) []entity.StaffMin {
	var staffs []entity.StaffMin
	db.connection.Model(&entity.Staff{}).Where(qstring).Offset((pagination["page"] - 1) * pagination["limit"]).Limit(pagination["limit"]).Find(&staffs)
	return staffs
}

func (db *konak) FindOneById(id int) entity.Staff {
	var oneuser entity.Staff
	db.connection.First(&oneuser, id)
	return oneuser
}

func (db *konak) FindOneByUsernameOrEmail(username string) entity.Staff {
	var oneuser entity.Staff
	err := 	db.connection.Where("username = ? OR email = ?", username, username).FirstOrInit(&oneuser).Error
	// db.connection.Where("username = ? OR email = ?", username, username).FirstOrInit(&oneuser)
	if err != nil {
		return oneuser
	}
	return oneuser
}

func (db *konak) FindByAccess(Access string) []entity.Staff {
	var allaccessuser []entity.Staff
	db.connection.Where("Access = ?", Access).Find(&allaccessuser)
	return allaccessuser
}

func (db *konak) FindByAccessCheck(id int) string {
	var akses entity.Staff
	db.connection.Where("id = ?", id).First(&akses)
	return akses.Access
}

func (db *konak) Update(staff entity.Staff) {
	db.connection.Updates(&staff)
}

func (db *konak) Delete(staff entity.Staff) {
	db.connection.Delete(&staff)
}
