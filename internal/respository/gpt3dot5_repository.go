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

type gpt3dot5Repository struct {
	AuthType    string
	ApiKey      string
	RequestType string
	Url         string
	ContentType string
	Model       string
	Db          *mongo.Collection
}

type gpt3dot5RequestBody struct {
	Model   string            `json:"model" bson:"model"`
	Message []messageTemplate `json:"messages" bson:"messages"`
}

type messageTemplate struct {
	Role    string `json:"role" bson:"role"`
	Content string `json:"content" bson:"content"`
}

type IGpt3dot5Repository interface {
	DoGpt3dot5AskQuestion(string) (*entity.Gpt3dot5Response, error)
	SaveMessage(*entity.Gpt3dot5Response, string, string) error
}

func NewGpt3dot5Repository(client *mongo.Client) *gpt3dot5Repository {
	repository := gpt3dot5Repository{}
	repository.AuthType = viper.GetString("Ia.Gpt3dot5.AuthType")
	repository.ApiKey = viper.GetString("Ia.Gpt3dot5.ApiKey")
	repository.ContentType = string("application/json")
	repository.RequestType = string("POST")
	repository.Url = viper.GetString("Ia.Gpt3dot5.Url")
	repository.Model = viper.GetString("Ia.Gpt3dot5.Model")

	collectionName := viper.GetString("MongoDb.Database.Collections.GptCollection")
	databaseName := viper.GetString("MongoDb.Database.Name")
	repository.Db = client.Database(databaseName).Collection(collectionName)

	return &repository
}

func (repo *gpt3dot5Repository) DoGpt3dot5AskQuestion(message string) (*entity.Gpt3dot5Response, error) {

	requestTemplate := gpt3dot5RequestBody{
		Model:   repo.Model,
		Message: []messageTemplate{{Role: "user", Content: message}},
	}

	requestBody, err := jsoniter.Marshal(requestTemplate)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(repo.RequestType, repo.Url, bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", repo.ContentType)

	auth := string(repo.AuthType + " " + repo.ApiKey)
	req.Header.Add("Authorization", auth)

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

	var result entity.Gpt3dot5Response
	err = jsoniter.Unmarshal(resultBody, &result)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &result, nil
}

func (r *gpt3dot5Repository) SaveMessage(data *entity.Gpt3dot5Response, id string, message string) error {
	insertData := entity.Gpt3dot5BackUpMessage{UserId: id, Data: *data, Prompt: message}
	_, err := r.Db.InsertOne(context.Background(), insertData)

	if err != nil {
		return err
	}

	return nil
}
