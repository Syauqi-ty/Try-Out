package entity

import "time"

type Disbursement struct {
	ID          string    `json:"id"`
	StaffID     string    `json:"staff_id"`
	ExternalID  string    `json:"external_id"`
	Amount      int       `json:"amount"`
	Description string    `json:"description" `
	BankCode    string    `json:"bank_code"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
