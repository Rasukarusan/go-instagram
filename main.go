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
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: true,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return origin == "http://localhost:9001"
		},
		AllowedMethods: []string{"POST"},
		AllowedHeaders: []string{
			"Accept", "Content-Type", "X-Custom-Header", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})
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
