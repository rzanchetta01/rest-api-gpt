package security

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

type responseBody struct {
	Results []struct {
		Flagged bool `json:"flagged" bson:"flagged"`
	} `json:"results" bson:"results"`
}

type requestTemplate struct {
	url         string
	apiKey      string
	contentType string
	authType    string
	requestType string
}
type requestBodyTemplate struct {
	Input string `json:"input" bson:"input"`
}

func CheckGPTCompliance(prompt string) bool {
	reqTemplate := requestTemplate{
		url:         viper.GetString("OpenAi.Moderation.Url"),
		apiKey:      viper.GetString("OpenAi.ApiKey"),
		contentType: "application/json",
		authType:    "Bearer",
		requestType: "POST",
	}

	reqBody, err := jsoniter.Marshal(requestBodyTemplate{Input: prompt})
	if err != nil {
		log.Println("ERROR CHECK GPT COMPLIANCE -> ", err)
		return false
	}

	req, err := http.NewRequest(reqTemplate.requestType, reqTemplate.url, bytes.NewBuffer(reqBody))
	req.Header.Add("Content-Type", reqTemplate.contentType)

	auth := string(reqTemplate.authType + " " + reqTemplate.apiKey)
	req.Header.Add("Authorization", auth)

	if err != nil {

		log.Println("ERROR CHECK GPT COMPLIANCE -> ", err)
		return false
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("ERROR CHECK GPT COMPLIANCE -> ", err)
		return false
	}

	resultBody, err := io.ReadAll(res.Body)
	if err != nil {

		log.Println("ERROR CHECK GPT COMPLIANCE -> ", err)
		return false
	}

	var result responseBody
	err = jsoniter.Unmarshal(resultBody, &result)
	if err != nil {
		log.Println("ERROR CHECK GPT COMPLIANCE -> ", err)
		return false
	}

	defer res.Body.Close()

	if result.Results[0].Flagged {
		log.Println("ERROR CHECK GPT COMPLIANCE -> ", errors.New("prompt got flagged = true"))
		return true
	}

	return false
}
