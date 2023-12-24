package models

type NewFavoriteParams struct {
	AccountID   string `json:"accountID"`
	FavoritedID string `json:"favoritedID"`
}
