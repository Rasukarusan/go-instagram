package main

import (
	"log"
	"net/http"
	"scraping/handlers"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	handler := handlers.Post{}
	router, err := rest.MakeRouter(
		rest.Post("/instagram", handler.List),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9000", api.MakeHandler()))
}
