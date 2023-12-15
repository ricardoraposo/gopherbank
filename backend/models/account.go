package models

import "time"

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
