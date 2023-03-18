package httputil

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/wilgun/joy-technologies-be/internal/constant"
	"github.com/wilgun/joy-technologies-be/internal/dto"
	"log"
	"net/http"
)

const (
	subjectRequest = "subject"
)

func GetBooksBySubjectRequest(r *http.Request, ps httprouter.Params) dto.UserGetBooksByGenreRequest {
	subject := ps.ByName(subjectRequest)

	return dto.UserGetBooksByGenreRequest{
		Subject: subject,
	}

}

func GetSubmitBookScheduleRequest(r *http.Request) (dto.SubmitBookScheduleRequest, error) {
	req := dto.SubmitBookScheduleRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("failed to decode submit book schedule request", err)
		return req, constant.ErrDecodeRequest
	}

	return req, nil
}

func GetAdminBooksBySubjectRequest(r *http.Request, ps httprouter.Params) dto.AdminGetBooksByGenreRequest {
	subject := ps.ByName(subjectRequest)

	return dto.AdminGetBooksByGenreRequest{
		Subject: subject,
	}

}
