package main

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

func NewPluginObject() interface{} {
	return Http{}
}

type Http struct{}

func (h Http) PluginObject() interface{} {
	return h
}

func (h Http) Request(method string, url string, body *string, headers map[string]string) (int, string, map[string]string, string) {
	respBody := ""

	var sendBody io.Reader
	sendBody = nil
	if body != nil {
		sendBody = strings.NewReader(*body)
	}
	req, err := http.NewRequest(method, url, sendBody)
	if err != nil {
		return 0, "", nil, err.Error()
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, "", nil, err.Error()
	}

	if resp.Body != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		respBodyStr := buf.String()
		respBody = respBodyStr
	}

	respHeaders := map[string]string{}
	for key, value := range resp.Header {
		respHeaders[key] = strings.Join(value, ",")
	}

	return resp.StatusCode, respBody, respHeaders, ""
}
