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
	Key      string    `json:"key"`
	UserId   int64     `json:"user_id"`
	BookTime time.Time `json:"book_time"`
}

type SubmitBookScheduleResponse struct {
	BookId              string    `json:"book_id"`
	ExpiredBookSchedule time.Time `json:"expired_book_schedule"`
}
