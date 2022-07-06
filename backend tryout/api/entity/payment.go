package entity

import (
	"time"
)

type Payment struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	ExternalID  string    `json:"external_id" gorm:"unique"`
	PayerID     uint64    `json:"payer_id"`
	Payer       Student   `json:"payer" gorm:"foreignKey:PayerID;references:ID"`
	TestID      uint64    `json:"test_id"`
	Test        Test      `json:"test" gorm:"foreignKey:TestID;references:ID;"`
	Amount      int       `json:"amount"`
	Method      string    `json:"method"`
	Status      string    `json:"status"`
	FailureCode string    `json:"failure_code,omitempty"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	FinishedAt  time.Time `json:"finishedAt" gorm:"column:finishedAt"`
}

type PaymentMin struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	ExternalID  string    `json:"external_id" gorm:"unique"`
	PayerID     uint64    `json:"payer_id"`
	Payer       Student   `json:"payer" gorm:"foreignKey:PayerID;references:ID"`
	TestID      uint64    `json:"test_id"`
	Test        TestBare  `json:"test" gorm:"foreignKey:TestID;references:ID;"`
	Amount      int       `json:"amount"`
	Method      string    `json:"method"`
	Status      string    `json:"status"`
	FailureCode string    `json:"failure_code,omitempty"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	FinishedAt  time.Time `json:"finishedAt" gorm:"column:finishedAt"`
}

type OVOCallback struct {
	ExternalID  string `json:"external_id" binding:"required"`
	Status      string `json:"status" binding:"required"`
	FailureCode string `json:"failure_code"`
}

type DanaCallback struct {
	ExternalID string `json:"external_id" binding:"required"`
	Status     string `json:"payment_status" binding:"required"`
}

type LinkAjaCallback struct {
	ExternalID string `json:"external_id" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

type CallbackHeader struct {
	XCallbackToken string `header:"x-callback-token"`
}

type PaymentRequest struct {
	StudentID uint64 `json:"student_id"`
	TestID    uint64 `json:"test_id" binding:"required"`
	Method    string `json:"method" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type PaymentResponse struct {
	ExternalID  string `json:"external_id"`
	CheckoutURL string `json:"checkout_url"`
	Amount      int    `json:"amount"`
	EwalletType string `json:"ewallet_type"`
}
