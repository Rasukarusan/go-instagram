package instagram

import (
	"encoding/json"
	"io"
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

/**
 * HTML内のJSONをgrepの要領で抜き出しており、io.Reader型ではなくJSON文字列で処理するため、
 * json.NewDecoderではなくjson.Unmarshalを使用する
 */
func (c Client) DecodeHTMLToJSON(html string) (*InstagramResponse, error) {
	r := regexp.MustCompile(`window._sharedData = {.*}`)
	jsonStr := strings.Replace(r.FindStringSubmatch(html)[0], "window._sharedData = ", "", 1)
	bytes := []byte(jsonStr)
	var response InstagramResponse
	err := json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}
