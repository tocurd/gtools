package gtools

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var Network networkInterface

type networkInterface interface {
	Get(string, map[string]string) (string, error)
	Post(url string, data string, headers map[string]string) (string, error)
}

type network struct{}

func init() {
	Network = &network{}
}

/**
 * @description: 发送GET请求
 * @param {string} url
 * @param {map[string]string} headers
 * @return {*}
 */
func (network) Get(url string, headers map[string]string) (string, error) {

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

/**
 * @description: post请求
 * @param {string} url
 * @param {map[string]string} data
 * @return {*}
 */
func (network) Post(url string, postData string, headers map[string]string) (string, error) {
	client := &http.Client{}

	request, err := http.NewRequest("POST", url, strings.NewReader(postData))
	if err != nil {
		return "", err
	}

	defer request.Body.Close()
	//增加header选项
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
