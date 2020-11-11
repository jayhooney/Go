package main

import (
	"fmt"
)


// 송신용 채널
func Out(outSignal chan<- string, signal string){
	outSignal <- signal
	fmt.Println(`OUT !`)
}

// 수신용 채널
func In(outSignal <-chan string, inSignal chan<- string){
	fmt.Println(`IN !`)
	signal := <- outSignal
	inSignal <- signal
	

}



func main() {
	fmt.Println(`===== CHANNEL START =====`)	

	inSignal := make(chan string,1)
	outSignal := make(chan string,1)
	
	go Out(outSignal, `FIGHTING ! ! ! `)

	go In(outSignal, inSignal)
	fmt.Println(`SIGNAL :`,<-inSignal)

	fmt.Println(`===== CHANNEL END   =====`)	
}


