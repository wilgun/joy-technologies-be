package webservice

import (
	"github.com/wilgun/joy-technologies-be/cmd/webservice/handler"
	"github.com/wilgun/joy-technologies-be/cmd/webservice/router"
	"github.com/wilgun/joy-technologies-be/internal/api/openlibrary"
	"github.com/wilgun/joy-technologies-be/internal/module"
	"github.com/wilgun/joy-technologies-be/internal/store"
	"log"
	"net/http"
)

const (
	httpPort = "8080"
)

func Start() {
	httpClient := &http.Client{}

	openLibraryClient := openlibrary.NewOpenLibaryApi(openlibrary.OpenLibraryParam{
		Client: httpClient,
	})

	bookStore := store.NewBookStore()

	bookModule := module.NewBookModule(module.BookModuleParam{
		OpenLibrary: openLibraryClient,
		BookStore:   bookStore,
	})

	httpHandler := handler.NewHttpHandler(handler.HttpHandlerImplParam{
		BookModule: bookModule,
	})

	r := router.Init(httpHandler)

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: r,
	}

	log.Printf("starting web, listening on %+v", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("error", err)
		log.Fatal("failed starting web on address", srv.Addr)
	}
}
