package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func StateReader(mutex *sync.Mutex, state map[int]int, readOps *uint64) {
	total := 0

	for {
		key := rand.Intn(5)

		// state 동기화를 위해 lock & Unlock
		mutex.Lock()
		total += state[key]
		mutex.Unlock()

		atomic.AddUint64(readOps, 1)
		time.Sleep(time.Millisecond)
	}
}

func StateWriter(mutex *sync.Mutex, state map[int]int, writeOps *uint64) {
	for {
		key := rand.Intn(5)
		val := rand.Intn(100)

		// state 동기화를 위해 lock & Unlock
		mutex.Lock()
		state[key] = val
		mutex.Unlock()

		atomic.AddUint64(writeOps, 1)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	fmt.Println(`===== MUTEX =====`)

	// 상태를 관리할 맵
	state := make(map[int]int)
	// 상태에 대한 접근을 동기화하는 뮤택스
	mutex := &sync.Mutex{}

	var readOps uint64 = 0
	var writeOps uint64 = 0

	// 100개의 고루틴으로 1 밀리초 마다 읽기 실행
	for readIdx, maxReadCnt := 0, 100; readIdx < maxReadCnt; readIdx++ {
		go StateReader(mutex, state, &readOps)
	}

	for writeIdx, maxWriteCnt := 0, 10; writeIdx < maxWriteCnt; writeIdx++ {
		go StateWriter(mutex, state, &writeOps)
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)

	fmt.Println(`readOps`, readOpsFinal)
	fmt.Println(`writeOps`, writeOpsFinal)

	mutex.Lock()
	fmt.Println(`state`, state)
	mutex.Unlock()

}
