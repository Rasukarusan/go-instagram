package instagram

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const baseURL = "https://www.instagram.com"

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
}

func NewClient() *Client {
	parsedURL, err := url.ParseRequestURI(baseURL)
	if err != nil {

	}
	return &Client{
		URL:        parsedURL,
		HTTPClient: http.DefaultClient,
	}
}

func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/html; charset=utf-8")
	return req, nil
}

type Post struct {
	Username string
	ImageURL string
	PostText string
	// OrgURL string
}

/**
 * InstagramにCurl→整形→JSONでレスポンス
 */
func (c *Client) GetPost(targetURL string) (*Post, error) {
	req, err := c.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	decoded, err := decode(res)
	if err != nil {
		return nil, err
	}

	return &Post{
		decoded.EntryData.PostPage[0].Graphql.ShortCodeMedia.DisplayURL,
		decoded.EntryData.PostPage[0].Graphql.ShortCodeMedia.Owner.Username,
		decoded.EntryData.PostPage[0].Graphql.ShortCodeMedia.EdgeMediaToCaption.Edges[0].Node.Text,
	}, nil
}

/**
 * HTML内のJSONをgrepの要領で抜き出しており、io.Reader型ではなくJSON文字列で処理するため、
 * json.NewDecoderではなくjson.Unmarshalを使用する
 */
func decode(resp *http.Response) (*InstagramResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html := string(body)

	r := regexp.MustCompile(`window._sharedData = {.*}`)
	jsonStr := strings.Replace(r.FindStringSubmatch(html)[0], "window._sharedData = ", "", 1)
	bytes := []byte(jsonStr)
	var response InstagramResponse
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}
