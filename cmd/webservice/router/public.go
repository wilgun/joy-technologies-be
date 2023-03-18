package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wilgun/joy-technologies-be/cmd/webservice/handler"
)

func publicRouter(router *httprouter.Router, handler handler.HttpHandler) {
	router.GET(
		"/public/v1/subjects/:subject",
		handler.GetBooksBySubject,
	)

	router.POST(
		"/public/v1/borrow",
		handler.SubmitBorrowBook,
	)
}
