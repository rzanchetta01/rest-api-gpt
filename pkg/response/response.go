package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseOKWithDataModel struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type responseOKModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type responseErrorModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type responseErrorCustomModel struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func ResponseOKWithData(c *gin.Context, data interface{}) {
	response := responseOKWithDataModel{
		Code:    200,
		Data:    data,
		Message: "OK",
	}

	c.JSON(http.StatusOK, response)
}

func ResponseCreated(c *gin.Context, data interface{}) {
	response := responseOKWithDataModel{
		Code:    201,
		Data:    data,
		Message: "Created",
	}

	c.JSON(http.StatusCreated, response)
}

func ResponseOK(c *gin.Context, message string) {
	response := responseOKModel{
		Code:    200,
		Message: message,
	}

	c.JSON(http.StatusOK, response)
}

func ResponseError(c *gin.Context, err string, code int) {
	response := responseErrorModel{
		Code:    99,
		Message: err,
	}

	c.JSON(code, response)
}

func ResponseCustomError(c *gin.Context, err interface{}, code int) {
	response := responseErrorCustomModel{
		Code:    99,
		Message: err,
	}

	c.JSON(code, response)
}
