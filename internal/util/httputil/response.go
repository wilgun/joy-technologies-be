package httputil

import (
	"encoding/json"
	"github.com/wilgun/joy-technologies-be/internal/dto"
	"net/http"
)

func writeFailedEncode(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(&dto.ResponseError{Message: "failed to encode response"})
}

func WriteErrorResponse(w http.ResponseWriter, rh dto.ResponseHandler) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rh.StatusCode)
	if err := json.NewEncoder(w).Encode(&dto.ResponseError{Message: rh.ErrorMessage, Code: rh.StatusCode}); err != nil {
		writeFailedEncode(w)
	}
}

func WriteSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		writeFailedEncode(w)
	}
}
