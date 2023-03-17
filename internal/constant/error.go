package constant

import "errors"

var (
	ErrInvalidSubject = errors.New("subject can't be an empty string")
)
