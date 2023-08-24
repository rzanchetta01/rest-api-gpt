package repository

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"project-p-back/internal/entity"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type gptDallERespository struct {
	AuthType    string
	ApiKey      string
	RequestType string
	Url         string
	ContentType string
	Db          *mongo.Collection
}

type gptDallERequestBody struct {
	Prompt string `json:"prompt,omitempty" bson:"prompt,omitempty"`
	N      int8   `json:"n,omitempty" bson:"n,omitempty"`
	Size   string `json:"size,omitempty" bson:"size,omitempty"`
}

type IGptDallERepository interface {
	DoGptDallEImageGeneration(string) (*entity.GptDallEResponse, error)
	GptDallESaveMessage(*entity.GptDallEResponse, string, string) error
}

func NewGptDallERepository(client *mongo.Client) *gptDallERespository {
	repository := gptDallERespository{}
	repository.AuthType = viper.GetString("Ia.DallE.AuthType")
	repository.ApiKey = viper.GetString("OpenAi.ApiKey")
	repository.Url = viper.GetString("Ia.DallE.Url")
	repository.ContentType = string("application/json")
	repository.RequestType = string("POST")

	collectionName := viper.GetString("MongoDb.Database.Collections.DallECollection")
	databaseName := viper.GetString("MongoDb.Database.Name")
	repository.Db = client.Database(databaseName).Collection(collectionName)

	return &repository
}

func (repo *gptDallERespository) DoGptDallEImageGeneration(message string) (*entity.GptDallEResponse, error) {
	requestTemplate := gptDallERequestBody{
		Prompt: message,
		N:      1,
		Size:   "256x256",
	}

	requestBody, err := jsoniter.Marshal(requestTemplate)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(repo.RequestType, repo.Url, bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", repo.ContentType)
	req.Header.Add("Authorization", repo.AuthType+" "+repo.ApiKey)

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resultBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result entity.GptDallEResponse
	err = jsoniter.Unmarshal(resultBody, &result)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &result, nil
}

func (repo *gptDallERespository) GptDallESaveMessage(data *entity.GptDallEResponse, userId string, message string) error {
	insertData := entity.GptDallEBackUpMessage{UserId: userId, Data: *data, Prompt: message}
	_, err := repo.Db.InsertOne(context.Background(), insertData)

	if err != nil {
		return err
	}

	return nil
}
