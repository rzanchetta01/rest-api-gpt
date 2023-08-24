package entity

import "github.com/graphql-go/graphql"

type GptDallERequest struct {
	Prompt string
	NImg   string
	Size   string
}

type GptDallEResponse struct {
	Created int64                  `json:"created,omitempty" bson:"created,omitempty"`
	Data    []gptDallEDataResponse `json:"data,omitempty" bson:"data,omitempty"`
}

type gptDallEDataResponse struct {
	Url string `json:"url,omitempty" bson:"url,omitempty"`
}

type GptDallEBackUpMessage struct {
	Id     string           `json:"id,omitempty" bson:"_id,omitempty"`
	UserId string           `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Data   GptDallEResponse `json:"data,omitempty" bson:"data,omitempty"`
	Prompt string           `json:"prompt,omitempty" bson:"prompt,omitempty"`
}

var gptDallEDataResponseGraphqlTemplate = graphql.NewObject(graphql.ObjectConfig{
	Name: "GptDallEDataResponse",
	Fields: graphql.Fields{
		"url": &graphql.Field{Type: graphql.String},
	},
})

var GptDallEResponseGraphqlTemplate = graphql.NewObject(graphql.ObjectConfig{
	Name: "GptDallEResponse",
	Fields: graphql.Fields{
		"created": &graphql.Field{Type: graphql.Int},
		"data":    &graphql.Field{Type: graphql.NewList(gptDallEDataResponseGraphqlTemplate)},
	},
})
