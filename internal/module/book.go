package module

import (
	"context"
	"fmt"
	"github.com/wilgun/joy-technologies-be/internal/api/openlibrary"
	"github.com/wilgun/joy-technologies-be/internal/constant"
	"github.com/wilgun/joy-technologies-be/internal/dto"
	"github.com/wilgun/joy-technologies-be/internal/model"
	"github.com/wilgun/joy-technologies-be/internal/store"
	"log"
	"time"
)

type BookWrapper interface {
	// User
	GetBooksBySubject(ctx context.Context, req dto.UserGetBooksByGenreRequest) (dto.UserGetBooksByGenreResponse, error)
	SubmitBookSchedule(ctx context.Context, req dto.SubmitBookScheduleRequest) (dto.SubmitBookScheduleResponse, error)

	// Librarian
	AdminGetBooksBySubject(ctx context.Context, req dto.AdminGetBooksByGenreRequest) (dto.AdminGetBooksByGenreResponse, error)
}

type bookModule struct {
	openLibrary openlibrary.Contract
	bookStore   store.BookStore
}

type BookModuleParam struct {
	OpenLibrary openlibrary.Contract
	BookStore   store.BookStore
}

func NewBookModule(param BookModuleParam) *bookModule {
	return &bookModule{
		openLibrary: param.OpenLibrary,
		bookStore:   param.BookStore,
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
	currentTime := time.Now().UTC()
	if len(req.Key) == 0 || req.UserId < 1 || req.BookTime.Before(currentTime) {
		return dto.SubmitBookScheduleResponse{}, constant.ErrInvalidSubmitSchedule
	}

	if diffDay := currentTime.Sub(req.BookTime).Hours() / 24; diffDay < -7 {
		return dto.SubmitBookScheduleResponse{}, constant.ErrInvalidSubmitSchedule
	}

	if b.bookStore.IsBookBorrowed(req.Key) {
		return dto.SubmitBookScheduleResponse{}, constant.ErrBookBorrowed
	}

	if user := b.bookStore.UserBorrowBook(req.UserId); user.UserId > 0 {
		return dto.SubmitBookScheduleResponse{}, constant.ErrUserBorrowingBook
	}

	if !b.eligibleSchedulePickupTime(req.BookTime) {
		return dto.SubmitBookScheduleResponse{}, constant.ErrNotEligiblePickUpTimeSchedule
	}

	borrowBook := b.bookStore.SubmitBorrowBook(model.UserBorrowBook{
		UserId:            req.UserId,
		ExpiredBorrowBook: req.BookTime.AddDate(constant.ExpiredBorrowYear, constant.ExpiredBorrowMonth, constant.ExpiredBorrowDay),
		BookKey:           req.Key,
	})

	schedule := b.bookStore.SubmitScheduleBook(borrowBook.BookId, req.BookTime)

	resp := dto.SubmitBookScheduleResponse{
		BookId:            borrowBook.BookId,
		StartPickUpBook:   schedule.StartPickUpBook,
		ExpiredPickUpBook: schedule.ExpiredPickUpBook,
	}

	return resp, nil
}

func (b *bookModule) eligibleSchedulePickupTime(bookTime time.Time) bool {
	schedulePickupTimeStart := bookTime.Add(-time.Minute * time.Duration(bookTime.Minute())).Add(-time.Second * time.Duration(bookTime.Second())).Add(-time.Nanosecond * time.Duration(bookTime.Nanosecond()))
	schedulePickupTimeEnd := schedulePickupTimeStart.Add(time.Hour * 1)

	key := fmt.Sprintf("%s-%s", schedulePickupTimeStart, schedulePickupTimeEnd)

	manyUser := b.bookStore.CheckManyUserAtTimeRange(key)

	if manyUser >= constant.MaxUserAtTimeRange {
		return false
	}

	return true
}

func (b *bookModule) AdminGetBooksBySubject(ctx context.Context, req dto.AdminGetBooksByGenreRequest) (dto.AdminGetBooksByGenreResponse, error) {
	if len(req.Subject) == 0 {
		return dto.AdminGetBooksByGenreResponse{}, constant.ErrInvalidSubject
	}

	books, err := b.openLibrary.GetBooksBySubject(ctx, openlibrary.UserGetBookRequest{
		Subject: req.Subject,
	})
	if err != nil {
		log.Printf("failed to get books from open library, err:%+v\n", err)
		return dto.AdminGetBooksByGenreResponse{}, constant.ErrGetBooksOpenLibrary
	}

	if len(books.Works) == 0 {
		log.Printf("books not found")
		return dto.AdminGetBooksByGenreResponse{}, constant.ErrBooksNotFound
	}

	//borrowedBooks := b.bookStore.GetListBorrowedBooks()
	//
	//booksData := []model.AdminBook{}
	//for _, work := range books.Works {
	//	authors := []string{}
	//	for _, author := range work.Authors {
	//		authors = append(authors, author.Name)
	//	}
	//
	//	userBook := model.UserBook{
	//		Key:           work.Key,
	//		Title:         work.Title,
	//		Author:        authors,
	//		EditionNumber: work.EditionCount,
	//	}
	//
	//	// TO DO: will be implemented
	//	//adminBook := model.AdminBook{
	//	//	UserBook: userBook,
	//	//}
	//	//for _, borrowedBook := range borrowedBooks {
	//	//	//if borrowedBook == userBook.Key {
	//	//	//	adminBook.PickUpSchedule.StartPickUpBook
	//	//	//}
	//	//}
	//
	//	//booksData = append(booksData, book)
	//}

	//respData := dto.AdminGetBooksByGenreResponse{Books: booksData}
	return dto.AdminGetBooksByGenreResponse{}, nil
}
