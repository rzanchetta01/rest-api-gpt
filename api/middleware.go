package api

import (
	"net/http"
	"project-p-back/pkg/jwtoken"
	"project-p-back/pkg/response"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Method", "POST, GET, DELETE, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func JWTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := jwtoken.TokenValid(ctx.Request)
		if err != nil {
			response.ResponseError(ctx, err.Error(), http.StatusUnauthorized)
			ctx.AbortWithStatus(401)
			return
		}

		ctx.Next()
	}
}
