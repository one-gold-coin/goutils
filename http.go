package goutils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type newHttp struct {
	url     string
	request interface{}
	http    *http.Request
}

func NewHttp(url, method string, request interface{}) (*newHttp, error) {
	h := &newHttp{
		url:     url,
		request: request,
	}
	requestJson, err := h.GetRequestJson()
	if err != nil {
		return nil, err
	}
	payload := strings.NewReader(requestJson)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	h.http = req
	return h, nil
}

func (h *newHttp) GetRequestJson() (string, error) {
	marshal, err := json.Marshal(h.request)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}

func (h *newHttp) SetHeader(key string, val string) *newHttp {
	h.http.Header.Add(key, val)
	return h
}

func (h *newHttp) Post(obj interface{}) error {
	client := &http.Client{
	}
	req := h.http
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(obj)
	if err != nil {
		return err
	}
	return nil
}

func (h *newHttp) Get(obj interface{}) error {
	client := &http.Client{
	}
	req := h.http
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(obj)
	if err != nil {
		return err
	}
	return nil
}
