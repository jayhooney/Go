package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var logger = fmt.Println

func goFunc(sigs chan os.Signal, done chan bool) {

	sig := <-sigs
	logger(``)
	logger(sig)
	done <- true
}

func main() {
	logger(`===== SIGNAL =====`)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// 지정한 시그널을 받을 수 있는 채널을 받고, 등록.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go goFunc(sigs, done)

	logger(`WAIT SIGNAL . . .`)
	<-done
	logger(`EXIT !`)

}
