package dto

type ResponseHandler struct {
	PublicResponse interface{}
	StatusCode     int
	ErrorMessage   string
}

type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
