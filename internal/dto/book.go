package dto

import "github.com/wilgun/joy-technologies-be/internal/model"

type UserGetBooksByGenreRequest struct {
	Subject string `json:"subject"`
}

type UserGetBooksByGenreResponse struct {
	Books []model.UserBook `json:"books"`
}
