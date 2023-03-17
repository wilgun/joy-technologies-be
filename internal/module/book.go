package module

import (
	"context"
	"github.com/wilgun/joy-technologies-be/internal/constant"
	"github.com/wilgun/joy-technologies-be/internal/dto"
)

type BookWrapper interface {
	GetBooksByGenre(ctx context.Context, req dto.UserGetBooksByGenreRequest) (dto.UserGetBooksByGenreResponse, error)
}

type BookModule struct {
}

func NewBookModule() *BookModule {
	return &BookModule{}
}

func (b *BookModule) GetBooksByGenre(ctx context.Context, req dto.UserGetBooksByGenreRequest) (dto.UserGetBooksByGenreResponse, error) {
	if len(req.Subject) == 0 {
		return dto.UserGetBooksByGenreResponse{}, constant.ErrInvalidSubject
	}

	return dto.UserGetBooksByGenreResponse{}, nil
}
