package store

import (
	"fmt"
	"github.com/wilgun/joy-technologies-be/internal/model"
	"math/rand"
	"strconv"
	"time"
)

var (
	ListScheduleBook   map[string][]string
	ListUserBorrowBook []model.UserBorrowBook
	ListBorrowedBook   []string
)

type BookStore interface {
	UserBorrowBook(userId int64) model.UserBorrowBook
	CheckManyUserAtTimeRange(key string) int
	SubmitBorrowBook(book model.UserBorrowBook) model.UserBorrowBook
	SubmitScheduleBook(bookId string, bookTime time.Time) model.ScheduleBook
	IsBookBorrowed(key string) bool
}

type bookStoreImpl struct {
}

func NewBookStore() BookStore {
	ListScheduleBook = map[string][]string{}
	return &bookStoreImpl{}
}

func (b *bookStoreImpl) UserBorrowBook(userId int64) model.UserBorrowBook {
	for _, user := range ListUserBorrowBook {
		if user.UserId == userId {
			return user
		}
	}
	return model.UserBorrowBook{}
}

func (b *bookStoreImpl) CheckManyUserAtTimeRange(key string) int {
	return len(ListScheduleBook[key])
}

func (b *bookStoreImpl) SubmitBorrowBook(book model.UserBorrowBook) model.UserBorrowBook {
	id := rand.Int()
	book.BookId = strconv.Itoa(id)
	ListUserBorrowBook = append(ListUserBorrowBook, book)
	ListBorrowedBook = append(ListBorrowedBook, book.BookKey)
	return book
}

func (b *bookStoreImpl) SubmitScheduleBook(bookId string, bookTime time.Time) model.ScheduleBook {
	schedulePickupTimeStart := bookTime.Add(-time.Minute * time.Duration(bookTime.Minute())).Add(-time.Second * time.Duration(bookTime.Second())).Add(-time.Nanosecond * time.Duration(bookTime.Nanosecond()))
	schedulePickupTimeEnd := schedulePickupTimeStart.Add(time.Hour * 1)

	key := fmt.Sprintf("%s-%s", schedulePickupTimeStart, schedulePickupTimeEnd)
	ListScheduleBook[key] = append(ListScheduleBook[key], bookId)

	return model.ScheduleBook{
		BookId:              bookId,
		ExpiredBookSchedule: schedulePickupTimeEnd,
	}
}

func (b *bookStoreImpl) IsBookBorrowed(key string) bool {
	for _, bookKey := range ListBorrowedBook {
		if key == bookKey {
			return true
		}
	}
	return false
}
