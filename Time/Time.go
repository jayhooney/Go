package main

import (
	"fmt"
	"time"
)

var logger = fmt.Println

func main() {

	logger(`===== Time =====`)
	// 현재 시간
	now := time.Now()
	logger(now)

	// 날짜 정보를 전달하여 time 구조체 생성
	myBirthday := time.Date(1993, 7, 31, 7, 31, 0, 731731, time.UTC)

	logger(myBirthday.Year())
	logger(myBirthday.Month())
	logger(myBirthday.Day())
	logger(myBirthday.Date())
	logger(myBirthday.Hour())
	logger(myBirthday.Minute())
	logger(myBirthday.Second())
	logger(myBirthday.Nanosecond())
	logger(myBirthday.Location())

	logger(myBirthday.Weekday())

	// 두 시간 비교 가능. 이전,이후,같은지 판단 가능.
	logger(myBirthday.Before(now))
	logger(myBirthday.After(now))
	logger(myBirthday.Equal(now))

	// 두 시간 사이의 간격 (Duraition)
	duration := now.Sub(myBirthday)
	logger(duration)

	logger(duration.Hours())
	logger(duration.Minutes())
	logger(duration.Seconds())
	logger(duration.Nanoseconds())

	// Add 로 시간을 더하거나, 줄일 수 있다.
	logger(myBirthday.Add(duration))
	logger(myBirthday.Add(-duration))

}
