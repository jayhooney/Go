package main

import (
	"fmt"
	"os"
)

var logger = fmt.Println

func main() {
	logger(`===== COMMANDLINE ARGS =====`)

	// 첫번째는 프로그램 경로, 1: 부터 프로그램 인자.
	argsWithProgram := os.Args
	argsWithoutProgram := os.Args[1:]

	// 특정 인자값 추출
	arg := os.Args[3]

	logger(argsWithProgram)
	logger(argsWithoutProgram)
	logger(arg)
}
