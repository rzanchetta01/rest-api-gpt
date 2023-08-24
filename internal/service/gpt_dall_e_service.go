package service

import (
	"errors"
	"project-p-back/internal/entity"
	repository "project-p-back/internal/respository"
	"project-p-back/pkg/security"

	"github.com/graphql-go/graphql"
)

type gptDallEService struct {
	repo repository.IGptDallERepository
}

type IGptDallEService interface {
	GptDallEGraphQlField() *graphql.Field
}

func NewGptDallEService(repo repository.IGptDallERepository) *gptDallEService {
	service := gptDallEService{}
	service.repo = repo

	return &service
}

func (service *gptDallEService) GptDallEGraphQlField() *graphql.Field {
	return &graphql.Field{
		Type: entity.GptDallEResponseGraphqlTemplate,
		Args: graphql.FieldConfigArgument{
			"message": &graphql.ArgumentConfig{Type: graphql.String},
			"user_id": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			message := p.Args["message"].(string)

			check := security.CheckGPTCompliance(message)
			if !check {
				return nil, errors.New("dall e prompt violates open ai compliance")
			}

			userId := p.Args["user_id"].(string)
			result, err := service.repo.DoGptDallEImageGeneration(message)
			if err != nil {
				return nil, err
			}

			if result.Data == nil {
				return nil, errors.New("response error, got empty data")
			}

			err = service.repo.GptDallESaveMessage(result, userId, message)
			if err != nil {
				return nil, err
			}

			return result, nil
		},
	}
}
