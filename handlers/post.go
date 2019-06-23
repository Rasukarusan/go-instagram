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
	resps := []*instagram.Result{}
	ch := make(chan *instagram.Result, len(param.URLs))
	for _, URL := range param.URLs {
		client := instagram.NewClient()
		// resp := client.GetResult(URL)
		go func(url string) {
			ch <- client.GetResult(url)
		}(URL)
	}
	for range param.URLs {
		resp := <-ch
		resps = append(resps, resp)
	}
	w.WriteJson(resps)
}
