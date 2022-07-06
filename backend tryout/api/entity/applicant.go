package entity

import "time"

type Applicant struct {
	//username,password,access update confusion
	ID          int    `json:"id" gorm:"primary_key;auto_increment"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	DivisionID  int    `json:"division_id"`
	Divisi Division `json:"divisi" gorm:"foreignKey:DivisionID;references:ID"`
	PositionID  int    `json:"position_id"`
	Position Posisi `json:"posisi" gorm:"foreignKey:PositionID;references:ID"`
	Name        string    `json:"name"`
	Motivasi 	string    `json:"motivasi"`
	Accepted    int    `json:"accepted" gorm:"default:0"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:createdAt"`
	ScheduledAt time.Time `json:"scheduled_at" gorm:"default:null"`
}
type ApplicantList struct {
	ID          int    `json:"id"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	DivisionID  int    `json:"division_id"`
	PositionID  int    `json:"position_id"`
	Name        string    `json:"name"`
	Motivasi 	string    `json:"motivasi"`
	Accepted    int    `json:"accepted"`
}
type ApplicantListWithDiv struct {
	ID          int    `json:"id"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Name        string    `json:"name"`
	Motivasi 	string    `json:"motivasi"`
	Accepted    int    `json:"accepted"`
	Divisiname string `json:"division_name"`
	Posisiname string `json:"posisi_name"`
}

