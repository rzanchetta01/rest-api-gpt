package config

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDb() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(viper.GetString("Mongodb.Database.ConnString"))

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	conStatus, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(conStatus, nil)
	if err != nil {
		return nil, err
	}

	log.Println("MongoDb sucessfully connected")
	return client, nil

}
