package main

import (
	"compress/flate"
	"crypto/tls"
	"fmt"
	"github.com/klauspost/compress/gzip"
	"io"
	"net/http"
	"time"
)

func curlGet(inputUrl, cookie string) string {
	tr := &http.Transport{
		IdleConnTimeout: 60 * time.Second,
		//Proxy: http.ProxyURL(&url.URL{
		//	Host: proxyInfo,
		//}),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 60,
	}
	request, err := http.NewRequest("GET", inputUrl, nil)
	request.Header.Add("Accept-Language", "zh-Hans-CN;q=1")
	request.Header.Add("Accept-Encoding", "gzip")
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.66 Safari/537.36 Edg/103.0.1264.44")
	request.Header.Add("cookie", cookie)
	if err != nil {
		return ""
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("response err")
		return ""
	}
	//defer response.Body.Close()
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)
	var body io.Reader
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		body, err = gzip.NewReader(response.Body)
		break
	case "deflate":
		body = flate.NewReader(response.Body)
		break
	default:
		body = response.Body
	}
	res, err := io.ReadAll(body)
	return string(res)
}
