package service

import (
	"errors"
	"project-p-back/internal/entity"
	repository "project-p-back/internal/respository"
	"project-p-back/pkg/security"

	"github.com/graphql-go/graphql"
)

type gpt3dot5Service struct {
	repo repository.IGpt3dot5Repository
}

type IGpt3dot5Service interface {
	Gpt3dot5GraphQlField() *graphql.Field
}

func NewGpt3dot5Service(repo repository.IGpt3dot5Repository) *gpt3dot5Service {
	var service = gpt3dot5Service{}
	service.repo = repo

	return &service
}

func (service *gpt3dot5Service) Gpt3dot5GraphQlField() *graphql.Field {
	return &graphql.Field{
		Type: entity.Gpt3dot5ResponseDataGraphqlTemplate,
		Args: graphql.FieldConfigArgument{
			"message": &graphql.ArgumentConfig{Type: graphql.String},
			"user_id": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			message := p.Args["message"].(string)

			isFlagged := security.CheckGPTCompliance(message)
			if isFlagged {
				return nil, errors.New("gpt3.5 prompt violates open ai compliance")
			}

			userId := p.Args["user_id"].(string)
			result, err := service.repo.DoGpt3dot5AskQuestion(message)
			if err != nil {
				return nil, err
			}

			if result.Id == "" {
				return nil, errors.New("response error, got empty fields")
			}

			err = service.repo.Gpt3dot5SaveMessage(result, userId, message)
			if err != nil {
				return nil, err
			}

			return result, nil
		},
	}
}
