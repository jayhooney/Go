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
	fmt.Println(`GO!`)


	msg <- `THE TIME IS NOW !`

}

func main() {
	fmt.Println(`===== CHANNEL SYNC START =====`)	

	msg := make(chan string, 1)

	go Logger(msg)

	fmt.Println(<-msg)


	fmt.Println(`===== CHANNEL SYNC END   =====`)	
}


