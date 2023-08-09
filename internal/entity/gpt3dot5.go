package entity

import "github.com/graphql-go/graphql"

type Gpt3dot5BackUpMessage struct {
	Id     string           `json:"id,omitempty" bson:"_id,omitempty"`
	UserId string           `json:"userId" bson:"userId"`
	Prompt string           `json:"promt" bson:"promt"`
	Data   Gpt3dot5Response `json:"data" bson:"data"`
}

type Gpt3dot5Response struct {
	Id      string               `json:"id" bson:"id"`
	Object  string               `json:"object" bson:"object"`
	Created int64                `json:"created" bson:"created"`
	Model   string               `json:"model" bson:"model"`
	Choices []gptResponseChoices `json:"choices" bson:"choices"`
	Usage   gptResponseTokens    `json:"usage" bson:"usage"`
}

type gptResponseChoices struct {
	Index        int                `json:"index" bson:"index"`
	Message      gptMessageTemplate `json:"message" bson:"message"`
	FinishReason string             `json:"finish_reason" bson:"finish_reason"`
}

type gptResponseTokens struct {
	PromptTokens     int `json:"prompt_tokens" bson:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens" bson:"completion_tokens"`
	TotalTokens      int `json:"total_tokens" bson:"total_tokens"`
}

type gptMessageTemplate struct {
	Role    string `json:"role" bson:"role"`
	Content string `json:"content" bson:"content"`
}

var gpt3dot5ResponseUsageGraphqlTemplate = graphql.NewObject(graphql.ObjectConfig{
	Name: "Gpt3dot5ResponseUsage",
	Fields: graphql.Fields{
		"prompt_tokens":     &graphql.Field{Type: graphql.Int},
		"completion_tokens": &graphql.Field{Type: graphql.Int},
		"total_tokens":      &graphql.Field{Type: graphql.Int},
	},
})

var gpt3dot5MessageTemplateGraphqlTemplate = graphql.NewObject(graphql.ObjectConfig{
	Name: "Gpt3dot5MessageTemplate",
	Fields: graphql.Fields{
		"role":    &graphql.Field{Type: graphql.String},
		"content": &graphql.Field{Type: graphql.String},
	},
})

var gpt3dot5ResponseChoicesGraphqlTemplate = graphql.NewObject(graphql.ObjectConfig{
	Name: "Gpt3dot5ResponseChoises",
	Fields: graphql.Fields{
		"index":         &graphql.Field{Type: graphql.Int},
		"message":       &graphql.Field{Type: gpt3dot5MessageTemplateGraphqlTemplate},
		"finish_reason": &graphql.Field{Type: graphql.String},
	},
})

var Gpt3dot5ResponseDataGraphqlTemplate = graphql.NewObject(graphql.ObjectConfig{
	Name: "Gpt3dot5ResponseData",
	Fields: graphql.Fields{
		"id":      &graphql.Field{Type: graphql.String},
		"object":  &graphql.Field{Type: graphql.String},
		"created": &graphql.Field{Type: graphql.Int},
		"model":   &graphql.Field{Type: graphql.String},
		"choices": &graphql.Field{Type: graphql.NewList(gpt3dot5ResponseChoicesGraphqlTemplate)},
		"usage":   &graphql.Field{Type: gpt3dot5ResponseUsageGraphqlTemplate},
	},
})
