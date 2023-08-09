package main

import (
	"fmt"
	"log"
	"project-p-back/api"
	"project-p-back/pkg/config"

	"github.com/spf13/viper"
)

func main() {
	config.SetupConfig("AppConfig")
	client, err := config.InitMongoDb()
	if err != nil {
		log.Fatalf(err.Error())
	}

	port := fmt.Sprintf(":%d", viper.GetInt("App.Port"))

	app := api.SetupRouter(client)
	app.Run(port)
}
