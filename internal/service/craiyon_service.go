package service

import (
	"errors"
	"project-p-back/internal/entity"
	repository "project-p-back/internal/respository"
)

type craiyonService struct {
	craiyonRepository repository.ICraiyonRepository
	userRepository    repository.IUserRepository
}

type ICraiyonService interface {
	GenerateCraiyonImages(string, string, string) (*entity.Craiyon, error)
}

func NewCraiyonService(craiyonRepo repository.ICraiyonRepository, userRepo repository.IUserRepository) *craiyonService {
	service := craiyonService{}
	service.craiyonRepository = craiyonRepo
	service.userRepository = userRepo

	return &service
}

func (service *craiyonService) GenerateCraiyonImages(prompt string, userId string, style string) (*entity.Craiyon, error) {

	_, err := service.userRepository.GetUserById(userId)

	if prompt == "" {
		return nil, errors.New("PROMPT MISSING")
	}

	if style == "" {
		return nil, errors.New("IMAGE STYLE MISSING")
	}

	if err != nil {
		return nil, errors.New("USER ID NOT EXISTS")
	}

	result, err := service.craiyonRepository.ExecCraiyonScript(prompt, userId, style)
	if err != nil {
		return nil, err
	}

	return result, nil
}
