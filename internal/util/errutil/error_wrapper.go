package errutil

import (
	"errors"
	"github.com/wilgun/joy-technologies-be/internal/constant"
	"github.com/wilgun/joy-technologies-be/internal/dto"
	"net/http"
)

func Wrap(err error) dto.ResponseHandler {
	httpCode := http.StatusInternalServerError
	switch {
	case errors.Is(err, constant.ErrMethodNotAllowed):
		httpCode = http.StatusMethodNotAllowed
	case errors.Is(err, constant.ErrRouterNotFound):
		httpCode = http.StatusNotFound
	case errors.Is(err, constant.ErrInvalidSubject) ||
		errors.Is(err, constant.ErrBooksNotFound) ||
		errors.Is(err, constant.ErrDecodeRequest) ||
		errors.Is(err, constant.ErrInvalidSubmitSchedule) ||
		errors.Is(err, constant.ErrNotEligiblePickUpTimeSchedule) ||
		errors.Is(err, constant.ErrBookBorrowed):
		httpCode = http.StatusBadRequest
	case errors.Is(err, constant.ErrUserBorrowingBook):
		httpCode = http.StatusNotAcceptable
	}

	return dto.ResponseHandler{
		StatusCode:   httpCode,
		ErrorMessage: err.Error(),
	}
}
