package main

import (
	"log"
	"net/http"
	"os"
	"scraping/handlers"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: true,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return origin == "https://rasukarusan.github.io/instagram-client"
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
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, api.MakeHandler()))
}
