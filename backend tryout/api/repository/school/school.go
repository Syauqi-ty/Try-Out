package repository

import (
	"gorm.io/gorm"
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
)

type SchoolRepo interface {
	FindSchoolBySlug(slug string) entity.School

	// BASIC CRUD
	CreateSchool(school entity.School) entity.School
	FindAllSchool() []entity.School
	FindSchoolByID(id int) entity.School
	UpdateSchool(school entity.School)
	DeleteSchool(school entity.School)
}

type schoolRepo struct {
	conn *gorm.DB
}

func NewSchoolRepo() SchoolRepo {
	conn := connection.Create()
	conn.AutoMigrate(&entity.School{})
	return &schoolRepo{conn}
}

func (r *schoolRepo) FindSchoolBySlug(slug string) entity.School {
	var school entity.School
	r.conn.Preload("Operators").Where("slug = ?", slug).First(&school)
	return school
}

////////////////
// BASIC CRUD //
////////////////

func (r *schoolRepo) CreateSchool(school entity.School) entity.School {
	r.conn.Create(school)
	return school
}

func (r *schoolRepo) FindAllSchool() []entity.School {
	var schools []entity.School
	r.conn.Preload("Operators").Find(&schools)
	return schools
}

func (r *schoolRepo) FindSchoolByID(id int) entity.School {
	var school entity.School
	r.conn.Preload("Operators").First(&school, id)
	return school
}

func (r *schoolRepo) UpdateSchool(school entity.School) {
	r.conn.Save(school)
}

func (r *schoolRepo) DeleteSchool(school entity.School) {
	r.conn.Delete(school)
}
