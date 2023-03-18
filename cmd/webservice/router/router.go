package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wilgun/joy-technologies-be/cmd/webservice/handler"
)

func Init(h *handler.HttpHandlerImpl) *httprouter.Router {
	router := httprouter.New()
	publicRouter(router, h)
	router.MethodNotAllowed = h.HandleMethodNotAllowed()
	router.NotFound = h.HandleMethodNotFound()

	return router
}
