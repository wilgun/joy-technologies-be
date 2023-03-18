package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wilgun/joy-technologies-be/cmd/webservice/handler"
)

func adminRouter(router *httprouter.Router, handler handler.HttpHandler) {
	router.GET(
		"/admin/v1/subjects/:subject",
		handler.AdminGetBooksBySubject,
	)
}
