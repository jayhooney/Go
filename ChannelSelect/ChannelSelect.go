package main

import (
	"fmt"
	"time"
)

func FuncA(chanA chan string) {
	time.Sleep(time.Second * 1)
	chanA <- "FIRST"
}

func FuncB(chanB chan string) {
	time.Sleep(time.Second * 5)
	chanB <- "SECOND"
}

func main() {
	chanA := make(chan string)
	chanB := make(chan string)

	go FuncA(chanA)

	go FuncB(chanB)

	for idx, len := 0, 2; idx < len; idx++ {
		// 다중 채널 연산을 select 로 기다리게 함.
		// 연산을 꼭 기다리고 한 번에 처리해야할 때 사용하나..?
		select {
		case msg1 := <-chanA:
			fmt.Println("received", msg1)
		case msg2 := <-chanB:
			fmt.Println("received", msg2)
		}
	}

}
