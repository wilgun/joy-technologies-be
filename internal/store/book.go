package store

import (
	"fmt"
	"github.com/wilgun/joy-technologies-be/internal/model"
	"math/rand"
	"time"
)

var (
	ListScheduleBook          map[string][]string
	ListUserBorrowBook        []model.UserBorrowBook
	ListBorrowedBooksSchedule map[string]model.ScheduleBook
)

type BookStore interface {
	UserBorrowBook(userId int64) model.UserBorrowBook
	CheckManyUserAtTimeRange(key string) int
	SubmitBorrowBook(book model.UserBorrowBook) model.UserBorrowBook
	SubmitScheduleBook(bookTime time.Time) model.ScheduleBook
	IsBookBorrowed(key string) bool
	GetListBorrowedBooksSchedule() map[string]model.ScheduleBook
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
	ListUserBorrowBook = append(ListUserBorrowBook, book)
	return book
}

func (b *bookStoreImpl) SubmitScheduleBook(bookTime time.Time) model.ScheduleBook {
	schedulePickupTimeStart := bookTime.Add(-time.Minute * time.Duration(bookTime.Minute())).Add(-time.Second * time.Duration(bookTime.Second())).Add(-time.Nanosecond * time.Duration(bookTime.Nanosecond()))
	schedulePickupTimeEnd := schedulePickupTimeStart.Add(time.Hour * 1)

	key := fmt.Sprintf("%s-%s", schedulePickupTimeStart, schedulePickupTimeEnd)
	id := rand.Int()
	scheduleId := string(id)
	ListScheduleBook[key] = append(ListScheduleBook[key], scheduleId)

	scheduleBook := model.ScheduleBook{
		ScheduleId:        scheduleId,
		StartPickUpBook:   schedulePickupTimeStart,
		ExpiredPickUpBook: schedulePickupTimeEnd,
	}

	ListBorrowedBooksSchedule[scheduleId] = scheduleBook

	return scheduleBook
}

func (b *bookStoreImpl) IsBookBorrowed(key string) bool {
	if _, ok := ListBorrowedBooksSchedule[key]; ok {
		return true
	}
	return false
}

func (b *bookStoreImpl) GetListBorrowedBooksSchedule() map[string]model.ScheduleBook {
	return ListBorrowedBooksSchedule
}
