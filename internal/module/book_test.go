package module

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/wilgun/joy-technologies-be/internal/api/openlibrary"
	"github.com/wilgun/joy-technologies-be/internal/constant"
	"github.com/wilgun/joy-technologies-be/internal/dto"
	"github.com/wilgun/joy-technologies-be/internal/model"
	"reflect"
	"testing"
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
