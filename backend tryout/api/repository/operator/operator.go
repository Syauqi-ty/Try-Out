package repository

import (
	conn "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
	"time"

	"gorm.io/gorm"
)

type OperatorRepo interface {
	FindOperatorWithFilter(params map[string]interface{}, pagination map[string]int) []entity.Operator
	FindAllOperator() []entity.Operator
	FindOperatorByID(id int) entity.Operator
	FindOperatorByUsername(username string) entity.Operator
	CreateOperator(operator entity.Operator) entity.Operator
	UpdateOperator(operator entity.Operator)
	DeleteOperator(operator entity.Operator)
}

type operatorRepo struct {
	conn *gorm.DB
}

func NewOperatorRepo() OperatorRepo {
	db := conn.Create()
	db.AutoMigrate(entity.Operator{})
	return &operatorRepo{db}
}

func (r *operatorRepo) FindOperatorWithFilter(params map[string]interface{}, pagination map[string]int) []entity.Operator {
	var operator []entity.Operator
	r.conn.Model(&entity.Operator{}).Where(params).Offset((pagination["page"] - 1) * pagination["limit"]).Limit(pagination["limit"]).Find(&operator)
	return operator
}

func (r *operatorRepo) FindOperatorByUsername(username string) entity.Operator {
	var operator entity.Operator
	r.conn.Model(&entity.Operator{}).Where("username = ? OR email = ?", username, username).First(&operator)
	return operator
}

func (r *operatorRepo) FindAllOperator() []entity.Operator {
	var operator []entity.Operator
	r.conn.Find(&operator)
	return operator
}

func (r *operatorRepo) FindOperatorByID(id int) entity.Operator {
	var operator entity.Operator
	r.conn.First(&operator, id)
	return operator
}

func (r *operatorRepo) CreateOperator(operator entity.Operator) entity.Operator {
	r.conn.Create(&operator)
	return operator
}

func (r *operatorRepo) UpdateOperator(operator entity.Operator) {
	operator.UpdatedAt = time.Now()
	r.conn.Model(&entity.Operator{}).Where("id = ?", int(operator.ID)).Updates(&operator)
}

func (r *operatorRepo) DeleteOperator(operator entity.Operator) {
	r.conn.Delete(&operator)
}
