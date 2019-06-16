package handlers

import (
	"net/http"
	"scraping/instagram"

	"github.com/ant0ine/go-json-rest/rest"
)

type Post struct{}

type postParam struct {
	URL string
}

/**
 * InstagramにCurl→整形→JSONでレスポンス
 */
func (p Post) List(w rest.ResponseWriter, req *rest.Request) {
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
