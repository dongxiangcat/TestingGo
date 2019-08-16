package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var timeout time.Duration

func init() {
	timeout = time.Duration(1000 * time.Millisecond) //超时时间1s
}

func Get(uri string, data map[string]interface{}) string {
	// 自定义一个Client
	client := &http.Client{Timeout: timeout}

	// 处理data参数
	var form url.Values = url.Values{}

	for key, value := range data {
		form.Add(key, ToString(value))
	}

	if strings.Contains(uri, "?") {
		uri += form.Encode()
	} else {
		uri += "?" + form.Encode()
	}

	resp, err := client.Get(uri)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return string(body)
	} else {
		return ""
	}

}

func Post(uri string, data map[string]interface{}) string {
	contentType := "application/x-www-form-urlencoded; charset=UTF-8"

	var form url.Values = url.Values{}
	for key, value := range data {
		form.Add(key, ToString(value))
	}

	formBody := bytes.NewBufferString(form.Encode())

	request, err := http.NewRequest("POST", uri, formBody) //请求
	request.Header.Set("Content-Type", contentType)

	http.DefaultClient.Timeout = timeout
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return ""
	}
	return string(body)
}
