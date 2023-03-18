package model

import "time"

type UserBook struct {
	Key           string   `json:"key"`
	Title         string   `json:"title"`
	Author        []string `json:"author"`
	EditionNumber int      `json:"edition_number"`
}

type ScheduleBook struct {
	BookId            string    `json:"book_id"`
	UserId            int64     `json:"user_id"`
	StartPickUpBook   time.Time `json:"start_pick_up_book"`
	ExpiredPickUpBook time.Time `json:"expired_book_schedule"`
}

type UserBorrowBook struct {
	BookId            string    `json:"book_id"`
	UserId            int64     `json:"user_id"`
	BookKey           string    `json:"book_key"`
	ExpiredBorrowBook time.Time `json:"expired_borrow_book"`
}
