package service

import (
	"github.com/dev-sota/di/repository"
)

type UserService interface {
	FindById(identifier int)
}

type userService struct {
	ur repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{ur}
}

func (us *userService) FindById(identifier int) {
	us.ur.FindById(identifier)
}
