package module

import (
	"context"
	"github.com/wilgun/joy-technologies-be/internal/dto"
)

type BookWrapper interface {
	GetBooksByGenre(ctx context.Context, req dto.UserGetBooksByGenreRequest)
}

type BookModule struct {
}

func NewBookModule() *BookModule {
	return &BookModule{}
}

func (b *BookModule) GetBooksByGenre(ctx context.Context, req dto.UserGetBooksByGenreRequest) dto.UserGetBooksByGenreResponse {

	return dto.UserGetBooksByGenreResponse{}
}