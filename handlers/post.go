package handlers

import (
	"net/http"
	"scraping/instagram"

	"github.com/ant0ine/go-json-rest/rest"
)

type Post struct{}

type postParam struct {
	URLs []string
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

	for _, URL := range param.URLs {
		client := instagram.NewClient()
		resp, err := client.GetResult(URL)
		if err != nil {
			return
		}
		w.WriteJson(resp)
	}
}
