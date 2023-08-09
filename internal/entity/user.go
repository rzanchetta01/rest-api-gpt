package entity

import (
	"fmt"
	"strings"
)

type User struct {
	Id       string `json:"id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

func (u *User) Validate(action string) error {
	var errMessages []string

	switch strings.ToLower(action) {
	//For each request type may have a different validation requirement
	default:
		if u.Username == "null" || u.Username == "" {
			errMessages = append(errMessages, "username cannot be empty or null")
		}
		if u.Password == "null" || u.Password == "" {
			errMessages = append(errMessages, "password cannot be empty or null")
		}
	}

	if len(errMessages) > 0 {
		return fmt.Errorf(strings.Join(errMessages, ", "))
	}

	return nil
}
