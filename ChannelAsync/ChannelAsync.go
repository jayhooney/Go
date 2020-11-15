package main

import (
	"fmt"
)

func main() {
	fmt.Println(`===== CHANNEL ASYNC =====`)

	messages := make(chan string)
	signals := make(chan bool)

	// messages 에서 값을 사용할 수 있는 경우엔 case절로 진입하지만
	// 그렇지 않은 경우엔 default 절로 진입.
	select {
	case msg := <-messages:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hello golang !"

	// Default 로 들어감
	select {
	case messages <- msg:
		fmt.Println(`sent message`, msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println(`received message`, msg)
	case sig := <-signals:
		fmt.Println(`recerived signal`, sig)
	default:
		fmt.Println(`nothing.`)
	}
}
