package constant

import "errors"

var (
	ErrMethodNotAllowed = errors.New("method not allowed")
	ErrRouterNotFound   = errors.New("router not found")

	ErrDecodeRequest = errors.New("failed to decode request")

	ErrInvalidSubject                = errors.New("subject can't be an empty string")
	ErrGetBooksOpenLibrary           = errors.New("failed to get books")
	ErrBooksNotFound                 = errors.New("books not found")
	ErrInvalidSubmitSchedule         = errors.New("invalid submit schedule request")
	ErrNotEligiblePickUpTimeSchedule = errors.New("there are many people at that time")
	ErrUserBorrowingBook             = errors.New("user is borrowing book")
)
