package repository

import (
	"time"
	"strconv"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"fmt"
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
	"bytes"
	"text/template"
	"net/smtp"
)
const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Study Buddy <id.studybuddy@gmail.com>"
const CONFIG_AUTH_EMAIL = "id.studybuddy@gmail.com"
const CONFIG_AUTH_PASSWORD = "obptdnvgiijdrjno"
var timezone string = viper.GetString("timezone")

type ApplicantRepo interface {
	FindAllApplicant() []entity.Applicant
	FindApplicantByID(id int) entity.Applicant
	CreateApplicant(applicant entity.Applicant) entity.Applicant
	UpdateApplicant(applicant entity.Applicant)
	DeleteApplicant(applicant entity.Applicant)
	FindOneByUsernameOrEmail(username string) entity.Applicant
	CreateSolution(solution entity.Solution) entity.Solution
	ListApplicant()[]entity.ApplicantList
	Update(ID int,applicant entity.Applicant)
	Email(applicant entity.Applicant)
}

type database struct {
	connection *gorm.DB
}

func NewApplicantRepo() ApplicantRepo {
	db := connection.Create()
	db.AutoMigrate(&entity.Applicant{})
	return &database{connection: db}
}

func (db *database) CreateApplicant(applicant entity.Applicant) entity.Applicant {
	loc, _ := time.LoadLocation(timezone)
	applicant.CreatedAt = time.Now().In(loc)
	db.connection.Create(&applicant)
	return applicant
}
func (db *database) ListApplicant()[]entity.ApplicantList {
	var applicant []entity.ApplicantList
	db.connection.Table("applicants").Find(&applicant)
	return applicant
}

func (db *database) FindAllApplicant() []entity.Applicant {
	var applicant []entity.Applicant
	db.connection.Find(&applicant)
	return applicant
}

func (db *database) FindApplicantByID(id int) entity.Applicant {
	var applicant entity.Applicant
	db.connection.First(&applicant, id)
	return applicant
}

func (db *database) UpdateApplicant(applicant entity.Applicant) {
	db.connection.Save(&applicant)
}

func (db *database) DeleteApplicant(applicant entity.Applicant) {
	db.connection.Delete(&applicant)
}

func (db *database) FindOneByUsernameOrEmail(username string) entity.Applicant {
	var oneuser entity.Applicant
	db.connection.Where("username = ? OR email = ?", username, username).First(&oneuser)
	return oneuser
}

func (db *database) CreateSolution(solution entity.Solution) entity.Solution {
	loc, _ := time.LoadLocation(timezone)
	solution.CreatedAt = time.Now().In(loc)
	solution.UpdatedAt = time.Now().In(loc)
	db.connection.Create(&solution)
	return solution
}

func (db *database) Update(ID int,applicant entity.Applicant){
	accepted := applicant.Accepted
	schedule := applicant.ScheduledAt
	db.connection.Model(&applicant).Where("id = ?", ID).Updates(entity.Applicant{Accepted:accepted,ScheduledAt:schedule})
}

func (db *database) Email(applicant entity.Applicant) {
	var user entity.Applicant
	email := applicant.Email
	db.connection.Where("email = ?",email).Find(&user)
	hari := user.ScheduledAt.Day()
	bulan := user.ScheduledAt.Month()
	tahun := user.ScheduledAt.Year()
	jam := user.ScheduledAt.Hour()
	menit := user.ScheduledAt.Minute()
	tanggal := strconv.Itoa(hari) + " " + bulan.String() + " " + strconv.Itoa(tahun) + " pukul " + strconv.Itoa(jam) + ":" + strconv.Itoa(menit)
	to := []string{email}
	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)
	lolos, _ := template.ParseFiles("templatelolos.html")
	gagal, _ := template.ParseFiles("templategagal.html")
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Selamat \n%s\n\n", mimeHeaders)))
	lolos.Execute(&body, struct {
		ScheduledAt    string
		Name    string
	  }{
		ScheduledAt:  tanggal,
		Name:    user.Name,
	  })
	var body2 bytes.Buffer
	mimeHeaders2 := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body2.Write([]byte(fmt.Sprintf("Subject: Mohon Maaf \n%s\n\n", mimeHeaders2)))
	gagal.Execute(&body2, struct {
		Name    string
	}{
		Name:    user.Name,
	})
	if user.Accepted == 1{
		smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, to, body.Bytes())
	}else if user.Accepted == 2{
		smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, to, body2.Bytes())
	}
}