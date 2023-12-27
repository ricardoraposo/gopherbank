package models

import (
	"time"
)

type Account struct {
	ID        int       `json:"-"`
	Number    string    `json:"number"`
	Password  string    `json:"password"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	Admin     bool      `json:"-"`
}

type NewAccountParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Number    string `json:"number"`
}

type LoginParams struct {
	Number   string `json:"number"`
	Password string `json:"password"`
}

type NewPasswordParams struct {
	Number   string `json:"number"`
	Password string `json:"password"`
}
