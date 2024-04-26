package gotify

import (
	"errors"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
)

var client = &http.Client{}

type Gotify struct {
	EndPoint string `json:"endPoint" yaml:"endPoint" comment:"服务器端点"`
	ApiKey   string `json:"apiKey" yaml:"apiKey" comment:"授权token"`
}

func NewGotify(apiKey, endPoint string) *Gotify {
	if strings.HasSuffix(endPoint, "/") {
		endPoint = endPoint[:len(endPoint)-1]
	}
	return &Gotify{EndPoint: endPoint, ApiKey: apiKey}
}

func (gotify *Gotify) do(req *http.Request) (resBody []byte, err error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gotify-Key", gotify.ApiKey)
	res, err := client.Do(req)
	if err != nil {
		return
	}
	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	errorDescription := gjson.GetBytes(resBody, "errorDescription")
	if errorDescription.Exists() {
		err = errors.New(errorDescription.String())
		return
	}
	return
}
