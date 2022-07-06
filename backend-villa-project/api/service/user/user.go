package service

import (
	"villa-akmali/api/helper/encrypt"
	"villa-akmali/api/model"
	repository "villa-akmali/api/repository/user"
)


type UserService interface {
	CreateUser(user model.User) model.User
	Login(user model.Login) model.User
}

type userService struct {
	userrepo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userService{
		userrepo: repo,
	}
}

func (service *userService) CreateUser(user model.User) model.User {
	enkrip := encrypt.Encrypt(user.Password)
	user.Password = enkrip
	return service.userrepo.CreateUser(user)
} 

func (service *userService) Login(user model.Login) model.User {
	return service.userrepo.Login(user)
} 