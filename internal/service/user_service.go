package service

import (
	"errors"
	"log"
	"project-p-back/internal/entity"
	repository "project-p-back/internal/respository"
	"project-p-back/pkg/security"
)

type userService struct {
	userRepository repository.IUserRepository
}

type IUserService interface {
	CreateUser(*entity.User) (*entity.User, error)
	LoginUser(*entity.User) (*entity.User, error)
}

func NewUserService(userRepo repository.IUserRepository) *userService {
	var userService = userService{}
	userService.userRepository = userRepo
	return &userService
}

func (service *userService) CreateUser(user *entity.User) (*entity.User, error) {
	isExists := service.userRepository.IsUserExists(user.Username)
	if isExists {
		return nil, errors.New("THIS USERNAME ALREADY EXISTS")
	}

	err := user.Validate("create")
	if err != nil {
		return nil, err
	}

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	result, err := service.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *userService) LoginUser(userAttempt *entity.User) (*entity.User, error) {
	err := userAttempt.Validate("login")
	if err != nil {
		return nil, err
	}

	result, err := service.userRepository.GetUserByUsername(userAttempt.Username)
	if err != nil {
		log.Println("LOGIN ERROR -> ", userAttempt, err)
		return nil, errors.New("login failed, incorrect username or password")
	}

	if security.VerifyPassword(result.Password, userAttempt.Password) == nil {
		return result, nil
	}

	return nil, errors.New("login failed, incorrect username or password")
}
