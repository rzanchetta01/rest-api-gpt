package entity

type Craiyon struct {
	Id              string   `json:"id,omitempty" bson:"_id,omitempty"`
	Images          []string `json:"images,omitempty" bson:"images,omitempty"`
	UserId          string   `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Prompt          string   `json:"prompt,omitempty" bson:"prompt,omitempty"`
	Style           string   `json:"style,omitempty" bson:"style,omitempty"`
	TemplateMessage string   `json:"template_message,omitempty" bson:"template_message,omitempty"`
}
