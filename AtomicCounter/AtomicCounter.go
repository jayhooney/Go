package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func Counter(counter *uint64) {
	for {
		atomic.AddUint64(counter, 1)

		time.Sleep(time.Millisecond)
	}
}

func main() {
	fmt.Println(`===== ATOMIC COUNTER =====`)

	// 항상 양수인 카운터를 나타내기 위해 부호 없는 정수를 사용.
	var ops uint64 = 0

	// 동시 업데이트 시뮬레이션.
	// 1밀리초 마다 카운터를 증가시키는 고루틴 50개를 띄움.
	for idx, len := 0, 50; idx < len; idx++ {
		go Counter(&ops)
	}

	// ops가 누적될 수 있도록 10초간 대기.
	time.Sleep(time.Second * 1)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops :", opsFinal)
}
