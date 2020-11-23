package main

import (
	"flag"
	"fmt"
)

var logger = fmt.Println

func main() {
	logger(`===== COMMANDLINE FLAG =====`)

	// 옵션값 - 기본값 - 설명 순서.
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("num", 731, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// 다른 곳에서 선언된 변수를 사용해 옵션을 선언하는것도 가능.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	logger("word", *wordPtr)
	logger("num", *numPtr)
	logger("fork", *boolPtr)
	logger("svar", svar)
	logger("tail", flag.Args())
}
