package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Account struct {
	ID        int       `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    string    `json:"number"`
	Password  string    `json:"password"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Admin     bool      `json:"-"`
}

type NewAccountParams struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
	Number    string `json:"number"`
}

func (a *NewAccountParams) Validate() error {
    validate := validator.New()
    return validate.Struct(a)
}
