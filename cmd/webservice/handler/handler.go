package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wilgun/joy-technologies-be/internal/module"
	"net/http"
)

type HttpHandler interface {
	//User
	GetBooksBySubject(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type HttpHandlerImpl struct {
	bookModule module.BookWrapper
}

type HttpHandlerImplParam struct {
	BookModule module.BookWrapper
}

func NewHttpHandler(param HttpHandlerImplParam) HttpHandler {
	return &HttpHandlerImpl{
		bookModule: param.BookModule,
	}
}
