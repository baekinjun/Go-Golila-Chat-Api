package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func HttpRequestPost(accessToken string, reqUrl string, data map[string]string) (result map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			return
		}
	}()
	method := http.MethodPost

	q := url.Values{}
	for k, v := range data {
		q.Add(k, v)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, reqUrl, strings.NewReader(q.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if accessToken != "" {
		req.Header.Set("Authorization", accessToken)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	response := &AIServerResponse{}

	if res.StatusCode == http.StatusOK {
		json.NewDecoder(res.Body).Decode(&response)
		if response.Code == 200 {
			result = response.Payload
		}
	} else {
		panic(fmt.Errorf("%s %s request failed", reqUrl, method))
	}

	return
}

func HttpRequestPut(accessToken string, reqUrl string, data map[string]string) (result map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			return
		}
	}()
	method := http.MethodPut

	client := &http.Client{}
	req, err := http.NewRequest(method, reqUrl, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if accessToken != "" {
		req.Header.Set("Authorization", accessToken)
	}
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	for k, v := range data {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	response := &AIServerResponse{}

	if res.StatusCode == http.StatusOK {
		json.NewDecoder(res.Body).Decode(&response)
		if response.Code == 200 {
			result = response.Payload
		}
	} else {
		panic(fmt.Errorf("%s %s request failed", reqUrl, method))
	}
	return
}

func HttpRequestGet(accessToken string, reqUrl string, params map[string]string) (result map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			return
		}
	}()
	method := http.MethodGet

	client := &http.Client{}
	req, err := http.NewRequest(method, reqUrl, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if accessToken != "" {
		req.Header.Set("Authorization", accessToken)
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	response := &AIServerResponse{}

	if res.StatusCode == http.StatusOK {
		json.NewDecoder(res.Body).Decode(&response)
		if response.Code == 200 {
			result = response.Payload
		}
	} else {
		panic(fmt.Errorf("%s %s request failed", reqUrl, method))
	}

	return
}

func HttpRequestDelete(accessToken string, reqUrl string, params map[string]string) (result map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			return
		}
	}()
	method := http.MethodDelete

	client := &http.Client{}
	req, err := http.NewRequest(method, reqUrl, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if accessToken != "" {
		req.Header.Set("Authorization", accessToken)
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	response := &AIServerResponse{}

	if res.StatusCode == http.StatusOK {
		json.NewDecoder(res.Body).Decode(&response)
		if response.Code == 200 {
			result = response.Payload
		}
	} else {
		panic(fmt.Errorf("%s %s request failed", reqUrl, method))
	}

	return
}
