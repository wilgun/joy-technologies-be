package module

import (
	"context"
	"fmt"
	"github.com/wilgun/joy-technologies-be/internal/api/openlibrary"
	"github.com/wilgun/joy-technologies-be/internal/constant"
	"github.com/wilgun/joy-technologies-be/internal/dto"
	"github.com/wilgun/joy-technologies-be/internal/model"
	"log"
	"time"
)

type BookWrapper interface {
	GetBooksBySubject(ctx context.Context, req dto.UserGetBooksByGenreRequest) (dto.UserGetBooksByGenreResponse, error)
	SubmitBookSchedule(ctx context.Context, req dto.SubmitBookScheduleRequest) (dto.SubmitBookScheduleResponse, error)
}

type bookModule struct {
	openLibrary openlibrary.Contract
}

type BookModuleParam struct {
	OpenLibrary openlibrary.Contract
}

func NewBookModule(param BookModuleParam) *bookModule {
	return &bookModule{
		openLibrary: param.OpenLibrary,
	}
}

func (b *bookModule) GetBooksBySubject(ctx context.Context, req dto.UserGetBooksByGenreRequest) (dto.UserGetBooksByGenreResponse, error) {
	if len(req.Subject) == 0 {
		return dto.UserGetBooksByGenreResponse{}, constant.ErrInvalidSubject
	}

	books, err := b.openLibrary.GetBooksBySubject(ctx, openlibrary.UserGetBookRequest{
		Subject: req.Subject,
	})
	if err != nil {
		log.Printf("failed to get books from open library, err:%+v\n", err)
		return dto.UserGetBooksByGenreResponse{}, constant.ErrGetBooksOpenLibrary
	}

	if len(books.Works) == 0 {
		log.Printf("books not found")
		return dto.UserGetBooksByGenreResponse{}, constant.ErrBooksNotFound
	}

	booksData := []model.UserBook{}

	for _, work := range books.Works {
		authors := []string{}
		for _, author := range work.Authors {
			authors = append(authors, author.Name)
		}

		book := model.UserBook{
			Key:           work.Key,
			Title:         work.Title,
			Author:        authors,
			EditionNumber: work.EditionCount,
		}

		booksData = append(booksData, book)
	}

	respData := dto.UserGetBooksByGenreResponse{Books: booksData}

	return respData, nil
}

func (b *bookModule) SubmitBookSchedule(ctx context.Context, req dto.SubmitBookScheduleRequest) (dto.SubmitBookScheduleResponse, error) {
	if len(req.Key) == 0 || req.UserId < 1 {
		return dto.SubmitBookScheduleResponse{}, constant.ErrInvalidSubmitSchedule
	}

	// TODO: will be implemented
	return dto.SubmitBookScheduleResponse{}, constant.ErrInvalidSubmitSchedule
}

func (b *bookModule) checkManyUserAtTimeRange(bookTime time.Time) bool {
	schedulePickupTimeStart := bookTime.Add(-time.Minute * time.Duration(bookTime.Minute())).Add(-time.Second * time.Duration(bookTime.Second())).Add(-time.Nanosecond * time.Duration(bookTime.Nanosecond()))
	schedulePickupTimeEnd := schedulePickupTimeStart.Add(time.Hour * 1)

	// TODO: will be continued when store already created
	fmt.Sprintf("%s-%s", schedulePickupTimeStart, schedulePickupTimeEnd)
	return true
}
