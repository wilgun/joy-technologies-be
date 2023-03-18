package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wilgun/joy-technologies-be/internal/constant"
	"github.com/wilgun/joy-technologies-be/internal/module"
	"github.com/wilgun/joy-technologies-be/internal/util/errutil"
	"github.com/wilgun/joy-technologies-be/internal/util/httputil"
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

func (h *HttpHandlerImpl) HandleMethodNotAllowed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := errutil.Wrap(constant.ErrMethodNotAllowed)
		httputil.WriteErrorResponse(w, result)
	}
}

func (h *HttpHandlerImpl) HandleMethodNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := errutil.Wrap(constant.ErrRouterNotFound)
		httputil.WriteErrorResponse(w, result)
	}
}
