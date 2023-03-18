package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/wilgun/joy-technologies-be/internal/util/httputil"
	"net/http"
)

func (h *HttpHandlerImpl) GetBooksBySubject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := httputil.GetBooksBySubjectRequest(r, ps)

	resp, err := h.bookModule.GetBooksBySubject(r.Context(), req)
	if err != nil {
		// TODO: will be implemented
		fmt.Println("")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
