package model

import "time"

type UserBook struct {
	BookId        string   `json:"book_id"`
	Title         string   `json:"title"`
	Author        []string `json:"author"`
	EditionNumber int      `json:"edition_number"`
}

type ScheduleBook struct {
	ScheduleId        string     `json:"schedule_id,omitempty"`
	UserId            int64      `json:"user_id,omitempty"`
	StartPickUpBook   *time.Time `json:"start_pick_up_book"`
	ExpiredPickUpBook *time.Time `json:"expired_book_schedule"`
}

type UserBorrowBook struct {
	ScheduleId        string    `json:"schedule_id"`
	UserId            int64     `json:"user_id"`
	BookId            string    `json:"book_id"`
	ExpiredBorrowBook time.Time `json:"expired_borrow_book"`
}

type AdminBook struct {
	UserBook
	PickUpSchedule ScheduleBook `json:"pick_up_schedule"`
}
