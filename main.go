package main

import (
    "instagram/api"
    "github.com/ant0ine/go-json-rest/rest"
    "log"
    "net/http"
    "net/url"
    "io/ioutil"
    "fmt"
)

type postParam struct {
    URL string 
}

type postResult struct {
    Username string
    ImageURL string
    PostText string
    // OrgURL string
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

    if param.URL  == "" {
        rest.Error(w, "URL is required", 400)
        return
    }

    resp,err := fetch(param.URL)
    if err != nil {
        fmt.Println(err)
        return
    }

    w.WriteJson(decode(resp))
}


/**
 * InstagramからHTMLBodyを取得
 */
func fetch(targetURL string) (*http.Response, error) {
    resp := &http.Response{}
    endpointURL, err := url.Parse(targetURL)
    if err != nil {
        return nil, err
    }
    resp,err = http.DefaultClient.Do(&http.Request{
        URL: endpointURL,
        Method: "GET",
        Header: http.Header {
            "Content-Type" : {"text/html; charset=utf-8"},
        },
    })
    return resp, err
}

func decode(resp *http.Response) postResult {
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    c := api.Client{}
    res, err := c.DecodeHTMLToJSON(string(body))
    if err != nil {
        fmt.Println(err)
    }

    return postResult{
        res.EntryData.PostPage[0].Graphql.ShortCodeMedia.DisplayURL,
        res.EntryData.PostPage[0].Graphql.ShortCodeMedia.Owner.Username,
        res.EntryData.PostPage[0].Graphql.ShortCodeMedia.EdgeMediaToCaption.Edges[0].Node.Text,
    }
}
