package model

import "time"

var (
	ListScheduleBook map[string][]string
)

type UserBook struct {
	Key           string   `json:"key"`
	Title         string   `json:"title"`
	Author        []string `json:"author"`
	EditionNumber int      `json:"edition_number"`
}

type ScheduleBook struct {
	BookId              string    `json:"book_id"`
	UserId              int64     `json:"user_id"`
	ExpiredBookSchedule time.Time `json:"expired_book_schedule"`
}
