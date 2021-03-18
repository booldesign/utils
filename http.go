package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 10:30
 * @Desc:
 */

//Deprecated
//请调用trace.HttpRequest(...)
func HttpRequest(queryUrl string, method string, params map[string]string, headers map[string]string) ([]byte, error) {
	method = strings.ToUpper(method)

	postParams := url.Values{}
	for k, v := range params {
		postParams.Set(k, v)
	}

	body := ioutil.NopCloser(strings.NewReader(postParams.Encode()))
	client := &http.Client{}
	client.Timeout = 30 * time.Second

	request, err := http.NewRequest(method, queryUrl, body)
	if err != nil {
		return []byte(""), err
	}

	if method == "POST" {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if len(headers) > 0 {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}

	//设置User-Agent
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:60.0) Gecko/20100101 Firefox/60.0")

	resp, err := client.Do(request)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return content, nil
}
