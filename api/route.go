package api

import (
	"project-p-back/internal/handler"
	repository "project-p-back/internal/respository"
	"project-p-back/internal/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	//repos
	userRepo := repository.NewUserRepository(client)
	gptRepo := repository.NewGpt3dot5Repository(client)

	//services
	userService := service.NewUserService(userRepo)
	gptService := service.NewGpt3dot5Service(gptRepo)

	//handlers
	userHandler := handler.NewUserHandler(userService)
	gptHandler := handler.NewGpt3dot5Handler(gptService)

	//routes
	userRoutes := router.Group("api/user")
	{
		userRoutes.POST("/", CORSMiddleware(), userHandler.CreateUser)
		userRoutes.POST("/login", CORSMiddleware(), userHandler.LoginUser)
	}

	graphqlRoutes := router.Group("api/graphql")
	{
		graphqlRoutes.POST("/gpt3dot5", JWTokenMiddleware(), CORSMiddleware(), gptHandler.Gpt3dot5ResolveQuery)
	}

	return router
}
