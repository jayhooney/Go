package main

import (
	"fmt"
	"os"
	"strings"
)

var logger = fmt.Println

func main() {
	logger(`===== ENVIRON =====`)

	// 환경변수의 k-v 를 설정
	os.Setenv("MY_NAME", "JAY")
	// 키에 매핑된 값 가져오기
	logger("MY_NAME", os.Getenv("MY_NAME"))
	logger("NODE_ENV", os.Getenv("NODE_ENV"))

	// 모든 환경 변수 목록
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		logger(pair[0])
	}

}
