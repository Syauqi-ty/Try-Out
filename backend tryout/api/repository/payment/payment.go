package repository

import (
	"gorm.io/gorm"
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
	preloader "studybuddy-backend-fast/api/helper/preloader"
)

type PaymentRepo interface {
	FindPaymentOfPayer(id int) []entity.Payment
	FindPaymentByExternalID(id string) entity.Payment
	UpdatePaymentByExternalID(externalID string, updatedData entity.Payment)

	FindAllPayment() []entity.Payment
	FindPaymentByID(id int) entity.Payment
	CreatePayment(payment entity.Payment) entity.Payment
	UpdatePayment(payment entity.Payment)
	DeletePayment(payment entity.Payment)
}

type paymentRepo struct {
	conn *gorm.DB
}

func NewPaymentRepo() PaymentRepo {
	db := connection.Create()
	db.AutoMigrate(&entity.Payment{})
	return &paymentRepo{db}
}

////////////////////////
// PRACTICAL USECASES //
////////////////////////

func (r *paymentRepo) FindPaymentOfPayer(id int) []entity.Payment {
	var payments []entity.Payment
	r.conn.Preload("Test", preloader.TestBarePreloader).Preload("Payer").Where("payer_id = ?", id).Find(&payments)
	return payments
}

func (r *paymentRepo) FindPaymentByExternalID(id string) entity.Payment {
	var payment entity.Payment
	r.conn.Preload("Test", preloader.TestBarePreloader).Preload("Payer").Where("external_id = ?", id).First(&payment)
	return payment
}

func (r *paymentRepo) UpdatePaymentByExternalID(externalID string, updatedData entity.Payment) {
	r.conn.Where("external_id = ?", externalID).Updates(updatedData)
}

////////////////
// BASIC CRUD //
////////////////

func (r *paymentRepo) FindAllPayment() []entity.Payment {
	var payments []entity.Payment
	r.conn.Preload("Test", preloader.TestBarePreloader).Preload("Payer").Find(&payments)
	return payments
}

func (r *paymentRepo) FindPaymentByID(id int) entity.Payment {
	var payment entity.Payment
	r.conn.Preload("Test", preloader.TestBarePreloader).Preload("Payer").First(&payment)
	return payment
}

func (r *paymentRepo) CreatePayment(payment entity.Payment) entity.Payment {
	r.conn.Create(&payment)
	return payment
}

func (r *paymentRepo) UpdatePayment(payment entity.Payment) {
	r.conn.Updates(&payment)
}

func (r *paymentRepo) DeletePayment(payment entity.Payment) {
	r.conn.Delete(&payment)
}
