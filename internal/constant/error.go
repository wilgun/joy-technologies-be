package constant

import "errors"

var (
	ErrMethodNotAllowed = errors.New("method not allowed")
	ErrRouterNotFound   = errors.New("router not found")

	ErrInvalidSubject      = errors.New("subject can't be an empty string")
	ErrGetBooksOpenLibrary = errors.New("failed to get books")
	ErrBooksNotFound       = errors.New("books not found")
)
