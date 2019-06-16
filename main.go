package main

import (
	"log"
	"net/http"
	"scraping/instagram"

	"github.com/ant0ine/go-json-rest/rest"
)

type postParam struct {
	URL string
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/instagram", postInstagram),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server started")
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9000", api.MakeHandler()))
}

/**
 * InstagramにCurl→整形→JSONでレスポンス
 */
func postInstagram(w rest.ResponseWriter, req *rest.Request) {
	param := postParam{}
	err := req.DecodeJsonPayload(&param)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if param.URL == "" {
		rest.Error(w, "URL is required", 400)
		return
	}

	client := instagram.NewClient()
	resp, err := client.GetPost(param.URL)

	w.WriteJson(resp)
}
