package main

import (
	"io/ioutil"
	"net/http"
)

type Network interface {
	Get(string, map[string]string) (string, error)
}

func Get(url string, headers map[string]string) (string, error) {

	client := http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	//增加header选项
	for key, value := range headers {
		reqest.Header.Add(key, value)
	}

	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	return string(body), nil
}
