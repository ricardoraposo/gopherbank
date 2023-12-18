package models

import (
	"time"
)

type Account struct {
	ID        int       `json:"-"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
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
	Number    string `json:"number"`
}
