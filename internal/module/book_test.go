package module

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/wilgun/joy-technologies-be/internal/api/openlibrary"
	"github.com/wilgun/joy-technologies-be/internal/constant"
	"github.com/wilgun/joy-technologies-be/internal/dto"
	"github.com/wilgun/joy-technologies-be/internal/model"
	"reflect"
	"testing"
	"time"
)

func TestGetBooksBySubject(t *testing.T) {
	mock := InitMock(t)

	type params struct {
		subject string
	}

	resp := dto.UserGetBooksByGenreResponse{
		Books: []model.UserBook{
			{
				Title:         "Book 1",
				Author:        []string{"author 1", "author 2"},
				EditionNumber: 1,
			},
			{
				Title:         "Book 2",
				Author:        []string{"author 3", "author 4"},
				EditionNumber: 2,
			},
		},
	}

	respOpenLibrary := openlibrary.UserGetBookResponse{
		Name:      "loves",
		WorkCount: 2,
		Works: []openlibrary.Work{
			{
				Title:        "Book 1",
				EditionCount: 1,
				Authors: []openlibrary.Author{
					{
						Name: "author 1",
					},
					{
						Name: "author 2",
					},
				},
			},
			{
				Title:        "Book 2",
				EditionCount: 2,
				Authors: []openlibrary.Author{
					{
						Name: "author 3",
					},
					{
						Name: "author 4",
					},
				},
			},
		},
	}
	tests := []struct {
		title          string
		params         params
		expectedError  error
		expectedResult dto.UserGetBooksByGenreResponse
		expectations   func(params *params)
	}{
		{
			title: "Success - Get Books By Subject",
			params: params{
				subject: "loves",
			},
			expectedError:  nil,
			expectedResult: resp,
			expectations: func(params *params) {
				mock.mockAPI.EXPECT().GetBooksBySubject(gomock.Any(), openlibrary.UserGetBookRequest{params.subject}).Return(respOpenLibrary, nil)
			},
		},
		{
			title: "Failed - Request Subject is an Empty String",
			params: params{
				subject: "",
			},
			expectedError:  constant.ErrInvalidSubject,
			expectedResult: dto.UserGetBooksByGenreResponse{},
			expectations: func(params *params) {
			},
		},
		{
			title: "Failed - Error Get Data From Open Library",
			params: params{
				subject: "loves",
			},
			expectedError:  constant.ErrGetBooksOpenLibrary,
			expectedResult: dto.UserGetBooksByGenreResponse{},
			expectations: func(params *params) {
				mock.mockAPI.EXPECT().GetBooksBySubject(gomock.Any(), openlibrary.UserGetBookRequest{params.subject}).Return(openlibrary.UserGetBookResponse{}, errors.New("error get books"))
			},
		},
		{
			title: "Success - Books Not Found",
			params: params{
				subject: "zxc",
			},
			expectedError:  constant.ErrBooksNotFound,
			expectedResult: dto.UserGetBooksByGenreResponse{},
			expectations: func(params *params) {
				mock.mockAPI.EXPECT().GetBooksBySubject(gomock.Any(), openlibrary.UserGetBookRequest{params.subject}).Return(openlibrary.UserGetBookResponse{}, nil)
			},
		},
	}

	for _, test := range tests {
		test.expectations(&test.params)
		result, err := mock.bookModule.GetBooksBySubject(context.Background(), dto.UserGetBooksByGenreRequest{
			Subject: test.params.subject,
		})

		if !errors.Is(err, test.expectedError) {
			t.Errorf("\ngot err  : %+v\nexpected : %+v", err, test.expectedError)
		}

		if !reflect.DeepEqual(result, test.expectedResult) {
			t.Errorf("got err: expected result: %+v, actual result: %+v", test.expectedResult, result)
		}
	}
}

func TestSubmitBookSchedule(t *testing.T) {
	mock := InitMock(t)

	bookTime := time.Now().UTC().AddDate(0, 0, 2)
	schedulePickupTimeStart := bookTime.Add(-time.Minute * time.Duration(bookTime.Minute())).Add(-time.Second * time.Duration(bookTime.Second())).Add(-time.Nanosecond * time.Duration(bookTime.Nanosecond()))
	schedulePickupTimeEnd := schedulePickupTimeStart.Add(time.Hour * 1)
	key := fmt.Sprintf("%s-%s", schedulePickupTimeStart, schedulePickupTimeEnd)

	bookTime7Days := time.Now().UTC().AddDate(0, 0, 7)
	schedulePickupTimeStart7Days := bookTime7Days.Add(-time.Minute * time.Duration(bookTime7Days.Minute())).Add(-time.Second * time.Duration(bookTime7Days.Second())).Add(-time.Nanosecond * time.Duration(bookTime7Days.Nanosecond()))
	schedulePickupTimeEnd7Days := schedulePickupTimeStart7Days.Add(time.Hour * 1)
	key7Days := fmt.Sprintf("%s-%s", schedulePickupTimeStart7Days, schedulePickupTimeEnd7Days)

	scheduleId := "123"
	bookTime8Days := time.Now().UTC().AddDate(0, 0, 8)
	expiredBorrowBook := bookTime.AddDate(constant.ExpiredBorrowYear, constant.ExpiredBorrowMonth, constant.ExpiredBorrowDay)
	expiredBorrowBook7Days := bookTime7Days.AddDate(constant.ExpiredBorrowYear, constant.ExpiredBorrowMonth, constant.ExpiredBorrowDay)

	type params struct {
		BookId   string
		UserId   int64
		BookTime time.Time
	}

	tests := []struct {
		title          string
		params         params
		expectedError  error
		expectedResult dto.SubmitBookScheduleResponse
		expectations   func(params *params)
	}{
		{
			title: "Success - Submit Book Schedule",
			params: params{
				BookId:   "asd",
				UserId:   1,
				BookTime: bookTime,
			},
			expectedError: nil,
			expectedResult: dto.SubmitBookScheduleResponse{
				BookId:            scheduleId,
				StartPickUpBook:   &schedulePickupTimeStart,
				ExpiredPickUpBook: &schedulePickupTimeEnd,
			},
			expectations: func(params *params) {
				mock.bookStore.EXPECT().IsBookBorrowed(params.BookId).Return(false)
				mock.bookStore.EXPECT().UserBorrowBook(params.UserId).Return(model.UserBorrowBook{})
				mock.bookStore.EXPECT().CheckManyUserAtTimeRange(key).Return(constant.MaxUserAtTimeRange - 1)
				mock.bookStore.EXPECT().SubmitBorrowBook(model.UserBorrowBook{
					ScheduleId:        scheduleId,
					UserId:            params.UserId,
					BookId:            params.BookId,
					ExpiredBorrowBook: params.BookTime.AddDate(constant.ExpiredBorrowYear, constant.ExpiredBorrowMonth, constant.ExpiredBorrowDay),
				}).Return(model.UserBorrowBook{
					ScheduleId:        scheduleId,
					UserId:            params.UserId,
					BookId:            params.BookId,
					ExpiredBorrowBook: expiredBorrowBook,
				})
				mock.bookStore.EXPECT().SubmitScheduleBook(params.BookTime).Return(model.ScheduleBook{
					ScheduleId:        scheduleId,
					StartPickUpBook:   &schedulePickupTimeStart,
					ExpiredPickUpBook: &schedulePickupTimeEnd,
				})
			},
		},
		{
			title: "Success - Submit Book Schedule One Week",
			params: params{
				BookId:   "asd",
				UserId:   1,
				BookTime: bookTime7Days,
			},
			expectedError: nil,
			expectedResult: dto.SubmitBookScheduleResponse{
				BookId:            scheduleId,
				StartPickUpBook:   &schedulePickupTimeStart7Days,
				ExpiredPickUpBook: &schedulePickupTimeEnd7Days,
			},
			expectations: func(params *params) {
				mock.bookStore.EXPECT().IsBookBorrowed(params.BookId).Return(false)
				mock.bookStore.EXPECT().UserBorrowBook(params.UserId).Return(model.UserBorrowBook{})
				mock.bookStore.EXPECT().CheckManyUserAtTimeRange(key7Days).Return(constant.MaxUserAtTimeRange - 1)
				mock.bookStore.EXPECT().SubmitBorrowBook(model.UserBorrowBook{
					ScheduleId:        scheduleId,
					UserId:            params.UserId,
					BookId:            params.BookId,
					ExpiredBorrowBook: params.BookTime.AddDate(constant.ExpiredBorrowYear, constant.ExpiredBorrowMonth, constant.ExpiredBorrowDay),
				}).Return(model.UserBorrowBook{
					ScheduleId:        scheduleId,
					UserId:            params.UserId,
					BookId:            params.BookId,
					ExpiredBorrowBook: expiredBorrowBook7Days,
				})
				mock.bookStore.EXPECT().SubmitScheduleBook(params.BookTime).Return(model.ScheduleBook{
					ScheduleId:        scheduleId,
					StartPickUpBook:   &schedulePickupTimeStart7Days,
					ExpiredPickUpBook: &schedulePickupTimeEnd7Days,
				})
			},
		},
		{
			title: "Failed - BookId not passed",
			params: params{
				BookId:   "",
				UserId:   1,
				BookTime: bookTime,
			},
			expectedError:  constant.ErrInvalidSubmitSchedule,
			expectedResult: dto.SubmitBookScheduleResponse{},
			expectations: func(params *params) {
			},
		},
		{
			title: "Failed - Userid not passed",
			params: params{
				BookId:   "asd",
				BookTime: bookTime,
			},
			expectedError:  constant.ErrInvalidSubmitSchedule,
			expectedResult: dto.SubmitBookScheduleResponse{},
			expectations: func(params *params) {
			},
		},
		{
			title: "Failed - Book time not passed",
			params: params{
				BookId: "asd",
				UserId: 1,
			},
			expectedError:  constant.ErrInvalidSubmitSchedule,
			expectedResult: dto.SubmitBookScheduleResponse{},
			expectations: func(params *params) {
			},
		},
		{
			title: "Failed - Pick up time more than 7 days",
			params: params{
				BookId:   "asd",
				UserId:   1,
				BookTime: bookTime8Days,
			},
			expectedError:  constant.ErrInvalidSubmitSchedule,
			expectedResult: dto.SubmitBookScheduleResponse{},
			expectations: func(params *params) {
			},
		},
		{
			title: "Failed - Book is being borrowed by another user",
			params: params{
				BookId:   "asd",
				UserId:   1,
				BookTime: bookTime,
			},
			expectedError:  constant.ErrBookBorrowed,
			expectedResult: dto.SubmitBookScheduleResponse{},
			expectations: func(params *params) {
				mock.bookStore.EXPECT().IsBookBorrowed(params.BookId).Return(true)

			},
		},
		{
			title: "Failed - User is borrowing book",
			params: params{
				BookId:   "asd",
				UserId:   1,
				BookTime: bookTime,
			},
			expectedError:  constant.ErrUserBorrowingBook,
			expectedResult: dto.SubmitBookScheduleResponse{},
			expectations: func(params *params) {
				mock.bookStore.EXPECT().IsBookBorrowed(params.BookId).Return(false)
				mock.bookStore.EXPECT().UserBorrowBook(params.UserId).Return(model.UserBorrowBook{
					UserId: params.UserId,
				})
			},
		},
		{
			title: "Failed - Many User at that time range",
			params: params{
				BookId:   "asd",
				UserId:   1,
				BookTime: bookTime,
			},
			expectedError:  constant.ErrNotEligiblePickUpTimeSchedule,
			expectedResult: dto.SubmitBookScheduleResponse{},
			expectations: func(params *params) {
				mock.bookStore.EXPECT().IsBookBorrowed(params.BookId).Return(false)
				mock.bookStore.EXPECT().UserBorrowBook(params.UserId).Return(model.UserBorrowBook{})
				mock.bookStore.EXPECT().CheckManyUserAtTimeRange(key).Return(constant.MaxUserAtTimeRange + 1)
			},
		},
	}

	for _, test := range tests {
		test.expectations(&test.params)
		result, err := mock.bookModule.SubmitBookSchedule(context.Background(), dto.SubmitBookScheduleRequest{
			BookId:   test.params.BookId,
			UserId:   test.params.UserId,
			BookTime: test.params.BookTime,
		})

		if !errors.Is(err, test.expectedError) {
			t.Errorf("\ngot err  : %+v\nexpected : %+v", err, test.expectedError)
		}

		if !reflect.DeepEqual(result, test.expectedResult) {
			t.Errorf("got err: expected result: %+v, actual result: %+v", test.expectedResult, result)
		}
	}
}

func TestAdminGetBooksBySubject(t *testing.T) {
	mock := InitMock(t)

	type params struct {
		subject string
	}

	currentTime := time.Now().UTC()
	startPickUpBook1 := currentTime
	expiredPickUpBook1 := currentTime.Add(time.Hour * 1)

	userBook1 := model.UserBook{
		BookId:        "111",
		Title:         "Book 1",
		Author:        []string{"author 1", "author 2"},
		EditionNumber: 1,
	}
	userBook2 := model.UserBook{
		BookId:        "222",
		Title:         "Book 2",
		Author:        []string{"author 3", "author 4"},
		EditionNumber: 2,
	}

	userBorrowBook := model.UserBorrowBook{
		ScheduleId: "123",
		UserId:     1,
		BookId:     "111",
	}

	borrowedBooks := map[string]model.UserBorrowBook{
		"111": userBorrowBook,
	}

	userBorrowBook2 := model.UserBorrowBook{
		ScheduleId: "456",
		UserId:     1,
		BookId:     "333",
	}

	borrowedBooks2 := map[string]model.UserBorrowBook{
		"456": userBorrowBook2,
	}

	scheduleBook := model.ScheduleBook{
		ScheduleId:        "123",
		UserId:            1,
		StartPickUpBook:   &startPickUpBook1,
		ExpiredPickUpBook: &expiredPickUpBook1,
	}

	borrowedBooksSchedule := map[string]model.ScheduleBook{
		"123": scheduleBook,
	}

	scheduleBook2 := model.ScheduleBook{
		ScheduleId:        "333",
		UserId:            1,
		StartPickUpBook:   &startPickUpBook1,
		ExpiredPickUpBook: &expiredPickUpBook1,
	}

	borrowedBooksSchedule2 := map[string]model.ScheduleBook{
		"333": scheduleBook2,
	}

	pickUpSchedule1 := model.ScheduleBook{
		StartPickUpBook:   &startPickUpBook1,
		ExpiredPickUpBook: &expiredPickUpBook1,
	}

	resp := dto.AdminGetBooksByGenreResponse{
		Books: []model.AdminBook{
			{
				UserBook:       userBook1,
				PickUpSchedule: pickUpSchedule1,
			},
			{
				UserBook: userBook2,
			},
		},
	}

	resp2 := dto.AdminGetBooksByGenreResponse{
		Books: []model.AdminBook{
			{
				UserBook: userBook1,
			},
			{
				UserBook: userBook2,
			},
		},
	}

	respOpenLibrary := openlibrary.UserGetBookResponse{
		Name:      "loves",
		WorkCount: 2,
		Works: []openlibrary.Work{
			{
				Key:          "111",
				Title:        "Book 1",
				EditionCount: 1,
				Authors: []openlibrary.Author{
					{
						Name: "author 1",
					},
					{
						Name: "author 2",
					},
				},
			},
			{
				Key:          "222",
				Title:        "Book 2",
				EditionCount: 2,
				Authors: []openlibrary.Author{
					{
						Name: "author 3",
					},
					{
						Name: "author 4",
					},
				},
			},
		},
	}
	tests := []struct {
		title          string
		params         params
		expectedError  error
		expectedResult dto.AdminGetBooksByGenreResponse
		expectations   func(params *params)
	}{
		{
			title: "Success - Get Books By Subject",
			params: params{
				subject: "loves",
			},
			expectedError:  nil,
			expectedResult: resp,
			expectations: func(params *params) {
				mock.mockAPI.EXPECT().GetBooksBySubject(gomock.Any(), openlibrary.UserGetBookRequest{params.subject}).Return(respOpenLibrary, nil)
				mock.bookStore.EXPECT().GetListBorrowedBook().Return(borrowedBooks)
				mock.bookStore.EXPECT().GetListBorrowedBooksSchedule().Return(borrowedBooksSchedule)
			},
		},
		{
			title: "Success - Get Books By Subject without pickup schedule",
			params: params{
				subject: "loves",
			},
			expectedError:  nil,
			expectedResult: resp2,
			expectations: func(params *params) {
				mock.mockAPI.EXPECT().GetBooksBySubject(gomock.Any(), openlibrary.UserGetBookRequest{params.subject}).Return(respOpenLibrary, nil)
				mock.bookStore.EXPECT().GetListBorrowedBook().Return(borrowedBooks2)
				mock.bookStore.EXPECT().GetListBorrowedBooksSchedule().Return(borrowedBooksSchedule2)
			},
		},
		{
			title: "Failed - Request Subject is an Empty String",
			params: params{
				subject: "",
			},
			expectedError:  constant.ErrInvalidSubject,
			expectedResult: dto.AdminGetBooksByGenreResponse{},
			expectations: func(params *params) {
			},
		},
		{
			title: "Failed - Error Get Data From Open Library",
			params: params{
				subject: "loves",
			},
			expectedError:  constant.ErrGetBooksOpenLibrary,
			expectedResult: dto.AdminGetBooksByGenreResponse{},
			expectations: func(params *params) {
				mock.mockAPI.EXPECT().GetBooksBySubject(gomock.Any(), openlibrary.UserGetBookRequest{params.subject}).Return(openlibrary.UserGetBookResponse{}, errors.New("error get books"))
			},
		},
		{
			title: "Success - Books Not Found",
			params: params{
				subject: "zxc",
			},
			expectedError:  constant.ErrBooksNotFound,
			expectedResult: dto.AdminGetBooksByGenreResponse{},
			expectations: func(params *params) {
				mock.mockAPI.EXPECT().GetBooksBySubject(gomock.Any(), openlibrary.UserGetBookRequest{params.subject}).Return(openlibrary.UserGetBookResponse{}, nil)
			},
		},
	}

	for _, test := range tests {
		test.expectations(&test.params)
		result, err := mock.bookModule.AdminGetBooksBySubject(context.Background(), dto.AdminGetBooksByGenreRequest{
			Subject: test.params.subject,
		})

		if !errors.Is(err, test.expectedError) {
			t.Errorf("\ngot err  : %+v\nexpected : %+v", err, test.expectedError)
		}

		if !reflect.DeepEqual(result, test.expectedResult) {
			t.Errorf("got err: expected result: %+v, actual result: %+v", test.expectedResult, result)
		}
	}
}
