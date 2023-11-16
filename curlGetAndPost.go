package main

import (
	"context"
	"fmt"
	"github.com/klauspost/compress/flate"
	"github.com/klauspost/compress/gzip"
	"golang.org/x/net/proxy"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

// curlGet  请求网址
func curlGet(inputUrl string) []byte {
	//tr := &http.Transport{
	//	IdleConnTimeout: 60 * time.Second,
	//	//Proxy: http.ProxyURL(&url.URL{
	//	//	Host: proxyInfo,
	//	//}),
	//	Proxy: http.ProxyURL(&url.URL{
	//		Host: "ip:port",
	//	}),
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// 跳过证书
	//}
	client := &http.Client{
		//Transport: tr,
		Timeout: time.Second * 30,
	}
	request, err := http.NewRequest("GET", inputUrl, nil)
	request.Header.Add("Accept-Language", "zh-Hans-CN;q=1")
	request.Header.Add("Accept-Encoding", "gzip")
	request.Header.Add("authority", "cloud.tencent.com")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.41")
	if err != nil {
		return nil
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("response err")
		return nil
	}
	//defer response.Body.Close()
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println("body close", err)
			os.Exit(2)
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
	return res
}

func curlGetSOCKS5(inputUrl string) []byte {
	dialer, err := proxy.SOCKS5("tcp", "192.168.0.66:7890", nil, proxy.Direct)
	if err != nil {
		fmt.Println(err, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	dialContext := func(ctx context.Context, network, address string) (net.Conn, error) {
		return dialer.Dial(network, address)
	}
	tr := &http.Transport{
		DialContext: dialContext,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 60,
	}
	request, err := http.NewRequest("GET", inputUrl, nil)
	request.Header.Add("Accept-Language", "zh-Hans-CN;q=1")
	request.Header.Add("Accept-Encoding", "gzip")
	request.Header.Add("Connection", "close")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.78")
	if err != nil {
		return nil
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("response err")
		fmt.Println(err)
		return nil
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
	return res
}
