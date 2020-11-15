package main

import "fmt"

func Receiver(jobs chan int, done chan bool) {
	for {
		j, more := <-jobs
		if more {
			fmt.Println("received job", j)
		} else {
			fmt.Println("received all jobs")
			done <- true
			return
		}
	}
}

func main() {

	fmt.Println(`===== Channel Close =====`)
	jobs := make(chan int, 100)
	done := make(chan bool)

	go Receiver(jobs, done)

	for idx, len := 0, 100; idx < len; idx++ {
		jobs <- idx
		fmt.Println("sent job", idx)
	}

	// len 만큼의 채널로 값을 보내고 난 뒤 채널 close
	close(jobs)

	fmt.Println("sent all jobs")

	// 리시버가 모든 jobs를 수신할 때 까지 대기.
	isBool := <-done

	fmt.Println("is it true?", isBool)
}
