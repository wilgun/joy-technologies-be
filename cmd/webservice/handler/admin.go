package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wilgun/joy-technologies-be/internal/util/errutil"
	"github.com/wilgun/joy-technologies-be/internal/util/httputil"
	"net/http"
)

func (h *HttpHandlerImpl) AdminGetBooksBySubject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := httputil.GetAdminBooksBySubjectRequest(r, ps)

	resp, err := h.bookModule.AdminGetBooksBySubject(r.Context(), req)
	if err != nil {
		result := errutil.Wrap(err)
		httputil.WriteErrorResponse(w, result)
		return
	}

	httputil.WriteSuccessResponse(w, resp)
}
