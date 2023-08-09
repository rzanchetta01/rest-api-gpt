package handler

import (
	"net/http"
	"project-p-back/internal/service"
	"project-p-back/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type gpt3dot5Handler struct {
	gptService service.IGpt3dot5Service
}

type queryRequest struct {
	Query string `json:"query"`
}

func NewGpt3dot5Handler(gpt3dot5Service service.IGpt3dot5Service) *gpt3dot5Handler {
	var handler = gpt3dot5Handler{}
	handler.gptService = gpt3dot5Service

	return &handler
}

func (handler *gpt3dot5Handler) Gpt3dot5ResolveQuery(c *gin.Context) {
	var reqQuery queryRequest
	err := c.ShouldBindJSON(&reqQuery)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Gpt3dot5",
		Fields: graphql.Fields{
			"GPT3dot5": handler.gptService.Gpt3dot5GraphQlField(),
		},
	})

	schemaConfig := graphql.SchemaConfig{Query: query}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	resulChan := make(chan graphql.Result)
	errChan := make(chan error)

	go func() {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: reqQuery.Query,
		})
		if len(result.Errors) > 0 {
			errChan <- result.Errors[0]
			return
		}

		resulChan <- *result
	}()

	select {
	case result := <-resulChan:
		response.ResponseOKWithData(c, &result)
		return
	case err := <-errChan:
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
}
