package handler

import (
	"net/http"
	"project-p-back/internal/entity"
	"project-p-back/internal/service"
	"project-p-back/pkg/response"

	"github.com/gin-gonic/gin"
)

type craiyonHandler struct {
	craiyonService service.ICraiyonService
}

type craiyonRequest struct {
	UserId string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Prompt string `json:"prompt,omitempty" bson:"prompt,omitempty"`
	Style  string `json:"style,omitempty" bson:"style,omitempty"`
}

func NewCraiyonHandler(craiyonService service.ICraiyonService) *craiyonHandler {
	handler := craiyonHandler{}
	handler.craiyonService = craiyonService

	return &handler
}

func (handler *craiyonHandler) CraiyonResolveRequest(c *gin.Context) {
	var req craiyonRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	resultChan := make(chan entity.Craiyon)
	errChan := make(chan error)

	go func() {
		result, err := handler.craiyonService.GenerateCraiyonImages(req.Prompt, req.UserId, req.Style)

		if err != nil {
			errChan <- err
			return
		}

		resultChan <- *result
	}()

	select {
	case result := <-resultChan:
		response.ResponseOKWithData(c, &result)
	case err := <-errChan:
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
	}

}
