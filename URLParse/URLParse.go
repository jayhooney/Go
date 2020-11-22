package main

import (
	"fmt"
	"net"
	"net/url"
)

var logger = fmt.Println

func main() {
	logger("===== URL PARSE =====")

	URL := "postgres://user:pass@host.com:5432/path?k=v&a=b&c=d&e=f#f"

	parsedURL, err := url.Parse(URL)

	if err != nil {
		panic(err)
	}

	// 스키마 접근
	logger(parsedURL.Scheme)

	// User 는 모든 인증정보를 가지고 있음.
	logger(parsedURL.User)
	logger(parsedURL.User.Username())
	pwd, _ := parsedURL.User.Password()
	logger(pwd)

	// Host 는 호스트명과 포트정보를 가지고 있음.
	logger(parsedURL.Host)
	host, port, _ := net.SplitHostPort(parsedURL.Host)
	logger(host)
	logger(port)

	// path , fragment 추출
	logger(parsedURL.Path)
	logger(parsedURL.Fragment)

	// 쿼리 파라미터 얻기
	logger(parsedURL.RawQuery)
	m, _ := url.ParseQuery(parsedURL.RawQuery)

	logger(m)
	logger(m["k"][0])
}
