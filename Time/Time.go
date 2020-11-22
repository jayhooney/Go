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

	logger(`===== TIME STAMP =====`)

	secs := now.Unix()
	nanos := now.UnixNano()
	logger(now)

	// UnixMillis 는 따로 없어서 아래 방법으로 계산해야함.
	millis := nanos / 1000000
	logger(secs)
	logger(millis)
	logger(nanos)

	logger(time.Unix(secs, 0))
	logger(time.Unix(0, nanos))

	logger(`===== TIME FORMATTING =====`)

	// RFC3339에 해당하는 레이아웃 상수를 사용해 시간을 포맷팅.
	logger(now.Format(time.RFC3339))

	timeA, _ := time.Parse(time.RFC3339, "2020-11-22T23:10:50+00:00")
	logger(timeA)

	// 커스텀 레이아웃
	logger(now.Format("7:31PM"))
	logger(now.Format("SUN NOV _2 06:31:31 2020"))
	logger(now.Format("2020-11-22T11:13:05.00.999999-07:00"))

	form := " 23 14 PM"

	timeB, _ := time.Parse(form, "7 31 PM")
	logger(timeB)

	// 이렇게 한땀 한땀 넣어줄 수도 있다.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	// 잘못된 입력값 들어오면 에러 반환
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	logger(e)

}
