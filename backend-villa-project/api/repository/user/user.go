package repository

import (
	"time"
	"villa-akmali/api/connection"
	"villa-akmali/api/helper/encrypt"
	"villa-akmali/api/model"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var timezone string = viper.GetString("timezone") 
type UserRepo interface {
	CreateUser(user model.User) model.User
	Login(login model.Login) model.User
	FindByID(id int) model.User
}

type database struct {
	connection *gorm.DB
}

func NewUserRepo() UserRepo {
	db := connection.Create()
	db.AutoMigrate(&model.User{})
	return &database{connection: db}
}

func (db *database) CreateUser(user model.User) model.User{
	var userkosong model.User
	loc,_ := time.LoadLocation(timezone)
	user.CreatedAt = time.Now().In(loc)
	bikinuser := db.connection.Create(&user)
	if bikinuser.Error != nil{
		return userkosong
	}else {
		return user
	}
}

func (db *database) Login(login model.Login) model.User  {
	var user model.User
	var userkosong model.User
	cariuser := db.connection.Table("users").Where("email=?",login.Email).First(&user)
	if cariuser.Error != nil{
		return userkosong
	} else {
		decrypt := encrypt.Decrypt(login.Password,user.Password)
		if decrypt == true{
			return user
		}else{
			userkosong.Name = "123456789"
			return userkosong
		}
	}
}

func (db *database) FindByID(id int) model.User {
	var user model.User
	if id == 0{
	}else{
		db.connection.Table("users").Where("id = ?",id).First(&user)
	}
	return user
}