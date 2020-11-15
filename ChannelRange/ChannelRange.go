package main

import (
	"fmt"
	"strconv"
)

func main() {

	messages := make(chan string, 1000)

	// messageQueue 크기 만큼 값 할당.
	for idx, len := 0, 1000; idx < len; idx++ {

		messages <- "THIS IS " + strconv.Itoa(idx) + " MESSAGE !"
		fmt.Println(strconv.Itoa(idx), `Message added !`)
	}

	// messages 채널 닫고
	close(messages)

	// 순서대로 메세지 출력
	for message := range messages {
		fmt.Println(message)
	}
}
