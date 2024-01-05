package models

type User struct {
	ID            int    `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	AccountNumber string `json:"accountNumber"`
}

type EditUserParams struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	PictureURL string `json:"pictureURL"`
}
