package repository

import (
	"fmt"
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
	"gorm.io/gorm"
	"strconv"
	"math/rand"
	"bytes"
	"text/template"
	"net/smtp"
)
const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Study Buddy <id.studybuddy@gmail.com>"
const CONFIG_AUTH_EMAIL = "id.studybuddy@gmail.com"
const CONFIG_AUTH_PASSWORD = "obptdnvgiijdrjno"

type StudentRepo interface {
	FindWithFilter(qstring map[string]interface{}, pagination map[string]int) []entity.Student
	FindAll() []entity.Student
	FindOneById(id int) entity.Student
	FindOneByUsernameOrEmail(username string) entity.Student
	CreateStudent(newUser entity.Student) entity.Student
	UpdateStudent(newUserData entity.Student)
	DeleteStudent(student entity.Student)
	Forget(student entity.Student) uint64
}
type database struct {
	connection *gorm.DB
}

func NewStudentRepo() StudentRepo {
	db := connection.Create()
	db.AutoMigrate(&entity.Student{})
	return &database{
		connection: db,
	}
}

func (db *database) FindAll() []entity.Student {
	var users []entity.Student
	db.connection.Find(&users)
	return users
}

func (db *database) FindWithFilter(qstring map[string]interface{}, pagination map[string]int) []entity.Student {
	var alluser []entity.Student
	db.connection.Where(qstring).Offset((pagination["page"] - 1) * pagination["limit"]).Limit(pagination["limit"]).Find(&alluser)
	return alluser
}
func RandomString(n int) string {
    var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
 
    s := make([]rune, n)
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
    return string(s)
}
func (db *database) FindOneById(id int) entity.Student {
	var oneuser entity.Student
	db.connection.First(&oneuser, id)
	return oneuser
}

func (db *database) FindOneByUsernameOrEmail(username string) entity.Student {
	var oneuser entity.Student
	db.connection.Where("username = ? OR email = ?", username, username).First(&oneuser)
	return oneuser
}

func (db *database) CreateStudent(newUser entity.Student) entity.Student {
	var hehe entity.Student
	data := db.connection.Where("username = ? OR email = ?", newUser.Username,newUser.Email).Find(&newUser)
	if data.RowsAffected != 0{
		return hehe
	} else {
		db.connection.Create(&newUser)
		return newUser
	}
}

func (db *database) UpdateStudent(newUserData entity.Student) {
	db.connection.Model(&entity.Student{}).Where("id = ?", int(newUserData.ID)).Updates(&newUserData)
}

func (db *database) DeleteStudent(student entity.Student) {
	db.connection.Unscoped().Delete(&student)
}
func (db *database) Forget(student entity.Student) uint64{
	var user entity.Student
	email := student.Email
	random := RandomString(6)
	random2 := RandomString(6)
	result := db.connection.Where("email = ?",email).Find(&user)
	id := strconv.FormatUint(user.ID,10)
	to := []string{email}
	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)
	t, _ := template.ParseFiles("templates.html")
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Reset Password \n%s\n\n", mimeHeaders)))
	t.Execute(&body, struct {
		Link    string
		Name    string
	  }{
		Link:    "https://studybuddy.id/reset/"+random+id+random2+"",
		Name:    user.Name,
	  })
    err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, to, body.Bytes())
    if err != nil {
    return 0
	} else if err == nil && result.RowsAffected == 0 {
	return 2
	}else{
	return 1
	}
}
