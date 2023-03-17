package module

import (
	"context"
	"github.com/wilgun/joy-technologies-be/internal/api/openlibrary"
	"github.com/wilgun/joy-technologies-be/internal/constant"
	"github.com/wilgun/joy-technologies-be/internal/dto"
	"github.com/wilgun/joy-technologies-be/internal/model"
	"log"
)

type BookWrapper interface {
	GetBooksBySubject(ctx context.Context, req dto.UserGetBooksByGenreRequest) (dto.UserGetBooksByGenreResponse, error)
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
			Title:         work.Title,
			Author:        authors,
			EditionNumber: work.EditionCount,
		}

		booksData = append(booksData, book)
	}

	respData := dto.UserGetBooksByGenreResponse{Books: booksData}

	return respData, nil
}
