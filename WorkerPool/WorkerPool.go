package main

import (
	"fmt"
	"strconv"
	"time"
)

func Worker(idx int, jobs <-chan int, result chan<- int) {

	for job := range jobs {
		workerName := `WORKER [` + strconv.Itoa(idx) + `]`
		fmt.Println(workerName, `START JOB`, job)
		time.Sleep(time.Second)
		fmt.Println(workerName, `FINISHED JOB`, job)

		result <- job * 10
	}

	return
}

func main() {

	fmt.Println("==== WORKER POOL =====")

	jobs, results := make(chan int, 100), make(chan int, 100)

	// worker 를 실행 시키고 잠시 블로킹
	for workerIdx, maxWorkerCnt := 1, 10; workerIdx < maxWorkerCnt; workerIdx++ {
		go Worker(workerIdx, jobs, results)
	}

	// 100개의 잡을 보내고
	for jobIdx, jobsLen := 0, 100; jobIdx < jobsLen; jobIdx++ {
		jobs <- jobIdx
	}

	// 더 이상 job이 없음을 알리기 위해 채널 close
	close(jobs)

	// 결과값들 가져오기
	for resultIdx, resultsLen := 0, 100; resultIdx < resultsLen; resultIdx++ {
		fmt.Println(`RESULT :`, strconv.Itoa(<-results))
	}

	return
}
