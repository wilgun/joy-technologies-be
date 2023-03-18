package httputil

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/wilgun/joy-technologies-be/internal/dto"
	"net/http"
)

const (
	subjectRequest = "subject"
)

func GetBooksBySubjectRequest(r *http.Request, ps httprouter.Params) dto.UserGetBooksByGenreRequest {
	param := r.Context().Value(httprouter.ParamsKey).(dto.UserGetBooksByGenreRequest)
	fmt.Println("param", param)

	subject := ps.ByName(subjectRequest)

	return dto.UserGetBooksByGenreRequest{
		Subject: subject,
	}

}
