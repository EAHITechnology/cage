package utils

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	PostMethod = "POST"
	GetMethod  = "GET"
)

/*
This Function will be replaced in the raptor erpc module later.
*/
func SendRequest(ctx context.Context, url string, body io.Reader, addHeaders map[string]string, method string) (resp []byte, err error) {
	fun := "sendRequest -->"
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("%s,%v", fun, err)
	}

	req.Header.Add("Content-Type", "application/json")
	if len(addHeaders) > 0 {
		for k, v := range addHeaders {
			req.Header.Add(k, v)
		}
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s,%v", fun, err)
	}

	if response == nil || response.Body == nil {
		return nil, fmt.Errorf("%s,%v", fun, "Body Nil")
	}

	defer response.Body.Close()

	respB, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("%s ReadAll err:%v", fun, err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%s http resp status = %d, %s", fun, response.StatusCode, string(respB))
	}

	return respB, nil
}
