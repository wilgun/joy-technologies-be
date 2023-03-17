package constant

import "errors"

var (
	ErrInvalidSubject      = errors.New("subject can't be an empty string")
	ErrGetBooksOpenLibrary = errors.New("failed to get books")
	ErrBooksNotFound       = errors.New("books not found")
)
