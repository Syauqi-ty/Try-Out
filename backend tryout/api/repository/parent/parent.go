package repository

import (
	"gorm.io/gorm"
	conn "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
	shahala "studybuddy-backend-fast/api/helper"
)

type ParentRepo interface {
	FindParentWithFilter(qstring map[string]interface{}, pagination map[string]int) []entity.Parent

	FindAllParent() []entity.Parent
	FindParentByID(id int) entity.Parent
	FindParentByUsername(username string) entity.Parent
	CreateParent(parent entity.Parent) entity.Parent
	UpdateParent(parent entity.Parent)
	DeleteParent(parent entity.Parent)
	Login(parent entity.Parent) uint64
}

type parentRepo struct {
	conn *gorm.DB
}

func NewParentRepo() ParentRepo {
	db := conn.Create()
	db.AutoMigrate(&entity.Parent{})
	return &parentRepo{db}
}

func (db *parentRepo) FindParentWithFilter(qstring map[string]interface{}, pagination map[string]int) []entity.Parent {
	var parent []entity.Parent
	db.conn.Model(&entity.Parent{}).Preload("Students").Where(qstring).Offset((pagination["page"] - 1) * pagination["limit"]).Limit(pagination["limit"]).Find(&parent)
	return parent
}

func (db *parentRepo) FindParentByUsername(username string) entity.Parent {
	var parent entity.Parent
	db.conn.Model(&entity.Parent{}).Where("username = ? OR email = ?", username, username).First(&parent)
	return parent
}

func (db *parentRepo) FindAllParent() []entity.Parent {
	var parent []entity.Parent
	db.conn.Preload("Students").Find(&parent)
	return parent
}

func (db *parentRepo) FindParentByID(id int) entity.Parent {
	var parent entity.Parent
	db.conn.Preload("Students").First(&parent, id)
	return parent
}

func (db *parentRepo) CreateParent(parent entity.Parent) entity.Parent {
	db.conn.Create(&parent)
	return parent
}

func (db *parentRepo) UpdateParent(parent entity.Parent) {
	db.conn.Model(&entity.Parent{}).Where("id = ?", int(parent.ID)).Updates(&parent)
}

func (db *parentRepo) DeleteParent(parent entity.Parent) {
	db.conn.Delete(&parent)
}

func (db *parentRepo) Login(parent entity.Parent) uint64{
	username := parent.Username
	password := parent.Password
	db.conn.Table("Parents").Where("email=? OR username=?",username,username).Find(&parent)
	id := parent.ID
	if shahala.Verify(password,parent.Password)==true{
		return id
	} else{
		return 0
	}
	
}
