package main

import (
	"fmt"
	"time"
)

func TimerExpirer(timer *time.Timer) {

	<-timer.C
	fmt.Println(`EXPIRED ! ! ! `)
}

func IsStop(isStop bool) {
	if isStop {
		fmt.Println("Timer stopped")
	}
}

func TickLogger(ticker *time.Ticker) {
	// 매 500ms마다 도착하는 값들을 순회하기 위해 채널에서 range를 사용합니다.
	for t := range ticker.C {
		fmt.Println(`Tick at`, t)
	}
}

func main() {
	fmt.Println(`===== TIMER AND TICKER =====`)
	// Timer는 미래 한 시점에 작업을 해야할 떄 사용
	// Ticker는 일정한 간격으로 작업을 반복해야할 때

	// 타이머가 만료될 때 까지 대기.
	timer := time.NewTimer(time.Second * 2)
	TimerExpirer(timer)

	// 타이머가 만료되기 전에 취소할 수 있다.
	timer = time.NewTimer(time.Second)
	go TimerExpirer(timer)

	// 값을 전송하는 채널을 사용.
	ticker := time.NewTicker(time.Millisecond * 500)
	// TickLogger(ticker)
	// Ticker 를 중간에 중단
	go TickLogger(ticker)

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println(`Ticker stopped`)

}
