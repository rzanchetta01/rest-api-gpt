package handler

import (
	"net/http"
	"project-p-back/internal/service"
	"project-p-back/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type gptHandler struct {
	gpt3dot5Service        service.IGpt3dot5Service
	gptDallEHandlerService service.IGptDallEService
}

type queryRequest struct {
	Query string `json:"query"`
}

func NewGptHandler(gpt3dot5Service service.IGpt3dot5Service, gptDallEService service.IGptDallEService) *gptHandler {
	return &gptHandler{
		gpt3dot5Service:        gpt3dot5Service,
		gptDallEHandlerService: gptDallEService,
	}
}

func (handler *gptHandler) GptResolveQuery(c *gin.Context) {
	var reqQuery queryRequest
	err := c.ShouldBindJSON(&reqQuery)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	baseQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "BaseQuery",
		Fields: graphql.Fields{
			"Gpt3dot5": handler.gpt3dot5Service.Gpt3dot5GraphQlField(),
			"GptDallE": handler.gptDallEHandlerService.GptDallEGraphQlField(),
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: baseQuery})
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	resultChan := make(chan graphql.Result)
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

		resultChan <- *result
	}()

	select {
	case result := <-resultChan:
		response.ResponseOKWithData(c, &result)
		return
	case err := <-errChan:
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
}
