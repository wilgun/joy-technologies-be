package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wilgun/joy-technologies-be/internal/util/errutil"
	"github.com/wilgun/joy-technologies-be/internal/util/httputil"
	"net/http"
)

func (h *HttpHandlerImpl) GetBooksBySubject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := httputil.GetBooksBySubjectRequest(r, ps)

	resp, err := h.bookModule.GetBooksBySubject(r.Context(), req)
	if err != nil {
		result := errutil.Wrap(err)
		httputil.WriteErrorResponse(w, result)
		return
	}

	httputil.WriteSuccessResponse(w, resp)

}

func (h *HttpHandlerImpl) SubmitBorrowBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req, err := httputil.GetSubmitBookScheduleRequest(r)
	if err != nil {
		result := errutil.Wrap(err)
		httputil.WriteErrorResponse(w, result)
		return
	}

	resp, err := h.bookModule.SubmitBookSchedule(r.Context(), req)
	if err != nil {
		result := errutil.Wrap(err)
		httputil.WriteErrorResponse(w, result)
		return
	}

	httputil.WriteSuccessResponse(w, resp)
}
