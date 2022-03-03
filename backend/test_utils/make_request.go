package test_utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

func MakeRequest(method, api string, param interface{}) (request *http.Request, err error) {
	method = strings.ToUpper(method)
	var (
		contentBuffer *bytes.Buffer
		jsonBytes     []byte
	)
	jsonBytes, err = json.Marshal(param)
	if err != nil {
		return
	}
	contentBuffer = bytes.NewBuffer(jsonBytes)
	request, err = http.NewRequest(string(method), api, contentBuffer)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	return
}
