package gotify

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (gotify Gotify) SendMessage(msg Message) (msgInfo MessageExternal, err error) {
	msgByte, err := json.Marshal(msg)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", gotify.EndPoint+"/message", bytes.NewReader(msgByte))
	if err != nil {
		return
	}
	body, err := gotify.do(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &msgInfo)
	return
}
