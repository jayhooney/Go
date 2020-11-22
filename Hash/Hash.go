package main

import (
	"crypto/md5"
	"fmt"
)

var logger = fmt.Println

func main() {
	logger(`===== HASH =====`)

	str := "my name is jay"

	hash := md5.New()

	hash.Write([]byte(str))

	logger(hash)

}
