package dtos

import "github.com/go-playground/validator/v10"

type Users struct {
	ID   uint   `json:"id" swaggerignore:"true"`
	Name string `json:"name" validate:"required,max=10"`
}

func (u Users) Validate() error {
	return validator.New().Struct(u)
}
