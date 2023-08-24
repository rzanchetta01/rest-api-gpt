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
	gpt3dot5Repo := repository.NewGpt3dot5Repository(client)
	gptDallERepo := repository.NewGptDallERepository(client)
	craiyonRepo := repository.NewCraiyonRepository(client)

	//services
	userService := service.NewUserService(userRepo)
	gpt3dot5Service := service.NewGpt3dot5Service(gpt3dot5Repo)
	gptDallEService := service.NewGptDallEService(gptDallERepo)
	craiyonService := service.NewCraiyonService(craiyonRepo, userRepo)

	//handlers
	userHandler := handler.NewUserHandler(userService)
	gptHandler := handler.NewGptHandler(gpt3dot5Service, gptDallEService)
	craiyonHandler := handler.NewCraiyonHandler(craiyonService)

	//routes
	userRoutes := router.Group("api/user")
	{
		userRoutes.POST("/", CORSMiddleware(), userHandler.CreateUser)
		userRoutes.POST("/login", CORSMiddleware(), userHandler.LoginUser)
	}

	webIaRoutes := router.Group("api/ia")
	{
		webIaRoutes.POST("/craiyon", CORSMiddleware(), JWTokenMiddleware(), craiyonHandler.CraiyonResolveRequest)
	}

	graphqlRoutes := router.Group("api/graphql")
	{
		graphqlRoutes.POST("/gpt", JWTokenMiddleware(), CORSMiddleware(), gptHandler.GptResolveQuery)
	}

	return router
}
