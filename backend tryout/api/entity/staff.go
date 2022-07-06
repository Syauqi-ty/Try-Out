package entity

import (
	"time"
)

type Staff struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Access    string    `json:"access"`
	Name      string    `json:"name"`
	Ma        string    `json:"ma"`
	Fi        string    `json:"fi"`
	Ki        string    `json:"ki"`
	Bi        string    `json:"bi"`
	Sos       string    `json:"sos"`
	Sej       string    `json:"sej"`
	Eko       string    `json:"eko"`
	Pu        string    `json:"pu"`
	Pk        string    `json:"pk"`
	Pmm       string    `json:"pmm"`
	Ppu       string    `json:"ppu"`
	Eng       string    `json:"eng"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type StaffMin struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Access    string    `json:"access"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s Staff) Min() StaffMin {
	return StaffMin{s.ID, s.Username, s.Password, s.Email, s.Phone, s.Access, s.Name, s.CreatedAt, s.UpdatedAt}
}
