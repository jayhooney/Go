package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// 리소스 이용을 제어하고 서비스의 품질을 유지하기 위한 중요 메커니즘
	fmt.Println(`===== RATE LIMIT =====`)

	requests := make(chan string, 5)

	fmt.Println(len(requests))

	// requests channel 에 임의값 보내기
	for reqIdx, reqLen := 0, 5; reqIdx < reqLen; reqIdx++ {
		requests <- `[` + strconv.Itoa(reqIdx) + `] REQUEST`
	}

	// 값을 다 보냈으므로 채널 닫기.
	close(requests)

	// 매 200 밀리초마다 값을 받아오는 속도 조절기.
	limter := time.Tick(time.Millisecond * 200)

	for reqeust := range requests {
		// request 를 수신하기 전에 limiter 채널 수신으로 블로킹하여 매 200밀리초마다 하나의 요청을 받아들이도록 속도 제한을 건다.
		<-limter
		fmt.Println(`RECEIVE REQUEST :`, reqeust, time.Now())
	}

	/*
		전반적으로는 속도제한을 유지하면서 요청들을 짧게 버스팅하고 싶은 경우
		limiter 채널을 버퍼링하여 해결할 수 있다.
	*/

	fmt.Println(`===== BURST =====`)

	// 최대 3개의 요청을 버스팅할 수 있는 채널 생성
	burstyLimiter := make(chan time.Time, 3)

	// 미리 버퍼 크기만큼 값 전달
	for burstyIdx, burstyLen := 0, 3; burstyIdx < burstyLen; burstyIdx++ {
		burstyLimiter <- time.Now()
	}

	// 버퍼가 비워지면 체워 넣을 고루틴
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyReqeusts := make(chan string, 5)
	for burstyReqIdx, burstyReqLen := 0, 5; burstyReqIdx < burstyReqLen; burstyReqIdx++ {
		burstyReqeusts <- `THIS IS [` + strconv.Itoa(burstyReqIdx) + `] BURSTY REQEUST !`
	}
	close(burstyReqeusts)

	for burstyReqeust := range burstyReqeusts {
		<-burstyLimiter
		fmt.Println(`RECEIVE REQUEST :`, burstyReqeust, time.Now())
	}

}
