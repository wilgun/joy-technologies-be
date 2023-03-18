package dto

import (
	"github.com/wilgun/joy-technologies-be/internal/model"
	"time"
)

type UserGetBooksByGenreRequest struct {
	Subject string `json:"subject"`
}

type UserGetBooksByGenreResponse struct {
	Books []model.UserBook `json:"books"`
}

type SubmitBookScheduleRequest struct {
	BookId   string    `json:"book_id"`
	UserId   int64     `json:"user_id"`
	BookTime time.Time `json:"book_time"`
}

type SubmitBookScheduleResponse struct {
	BookId            string     `json:"book_id"`
	StartPickUpBook   *time.Time `json:"start_pick_up_book"`
	ExpiredPickUpBook *time.Time `json:"expired_pick_up_book"`
}

type AdminGetBooksByGenreRequest struct {
	Subject string `json:"subject"`
}

type AdminGetBooksByGenreResponse struct {
	Books []model.AdminBook `json:"books"`
}
