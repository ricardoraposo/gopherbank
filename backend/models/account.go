package models

import "time"

type Account struct {
	ID        int       `json:"id,omitempty"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Admin     bool      `json:"admin"`
}
