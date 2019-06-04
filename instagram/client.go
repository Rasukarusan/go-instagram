package api

import (
	"encoding/json"
	"regexp"
	"strings"
)

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

/**
 * @TODO 空実装
 */
type Client struct{}

func (c Client) NewClient()  {}
func (c Client) NewRequest() {}
