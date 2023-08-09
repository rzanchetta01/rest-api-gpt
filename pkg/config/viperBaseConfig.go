package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func SetupConfig(configFileName string) {
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yml")

	root, _ := os.Getwd()

	viper.AddConfigPath(filepath.Join(root, "/configurations"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("error reading config -> ", err)
	}
}
