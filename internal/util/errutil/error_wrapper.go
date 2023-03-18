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
	case errors.Is(err, constant.ErrInvalidSubject) || errors.Is(err, constant.ErrBooksNotFound):
		httpCode = http.StatusBadRequest
	}

	return dto.ResponseHandler{
		StatusCode:   httpCode,
		ErrorMessage: err.Error(),
	}
}
