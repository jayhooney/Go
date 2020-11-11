package main

import (
	"time"
	"fmt"
)


func Logger(msg chan string){

	
	fmt.Println(`WAIT WAIT WAIT`)
	time.Sleep(3 * time.Second)
	fmt.Println(`ALRIGHT!`)
	time.Sleep(2 * time.Second)
	fmt.Println(`THE TIME IS NOW !`)


	msg <- `GO !`

}

func main() {
	fmt.Println(`===== CHANNEL SYNC START =====`)	

	msg := make(chan string, 1)

	go Logger(msg)


	// <- msg
	// msg 채널의 값을 강제로 받아오게끔 해서 동기화 시키는 방법.
	fmt.Println(<-msg)


	fmt.Println(`===== CHANNEL SYNC END   =====`)	
}


