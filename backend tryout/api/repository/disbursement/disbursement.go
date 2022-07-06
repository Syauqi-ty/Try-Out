package repository

import (
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"

	"github.com/spf13/viper"
	"github.com/xendit/xendit-go"
	"gorm.io/gorm"
)

var timezone string = viper.GetString("timezone")

type DisbursementRepo interface {
	FindAllDisbursement() []entity.Disbursement
	CreateDisbursement(disbursement entity.Disbursement) entity.Disbursement
}

type database struct {
	connection   *gorm.DB
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

func NewDisbursementRepo() DisbursementRepo {
	db := connection.Create()
	db.AutoMigrate(&entity.Disbursement{})
	return &database{connection: db}
}

func (db *database) CreateDisbursement(disbursement entity.Disbursement) entity.Disbursement {
	db.connection.Create(&disbursement)
	return disbursement
}

func (db *database) FindAllDisbursement() []entity.Disbursement {
	var disbursement []entity.Disbursement
	db.connection.Find(&disbursement)
	return disbursement
}
