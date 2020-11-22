package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var logger = fmt.Println

func main() {
	logger(`===== REGEX =====`)

	// 패턴과 문자열이 일치하는지 여부 판단.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	logger(match)

	// 위에서는 바로 매치했지만, 다른 정규식 작업에서는 최적화된 Regexp 구조체를 compile 해야함.
	rgxCompile, _ := regexp.Compile("p([a-z]+)ch")

	logger(rgxCompile.MatchString("peach"))

	// 정규식과 일치하는 문자열 찾기
	logger(rgxCompile.FindString("peach peach apeach"))

	// 첫번째로 매칭되는 문자열을 찾아 첫,마지막 인덱스를 반환.
	logger(rgxCompile.FindStringIndex("peach peach apeach"))

	// 전체 패턴 일치와 해당 일치의 부분 일치에 대한 정보를 모두 포함.
	logger(rgxCompile.FindStringSubmatch("peach peach apeach"))

	// 전체 일치와 부분 일치의 인덱스에 대한 정보 반환.
	logger(rgxCompile.FindStringSubmatchIndex("peach peach apeach"))

	//모든 일치 작업에 적용.
	// 두번째 인자로 양수를 넘겨주면 일치 항목의 개수를 제한 할수 있다.
	logger(rgxCompile.FindAllString("peach peach apeach", -1))
	logger(rgxCompile.FindAllString("peach peach apeach", 2))

	// 다른 함수에 대해서도 사용가능.
	logger(rgxCompile.FindAllStringSubmatchIndex("peach peach apeach", -1))

	// string 말고 []byte로 인자를 넘겨 줄수도 있다.
	logger(rgxCompile.Match([]byte("peach")))

	// 정규표현식으로 상수를 만들때 Complie의 변형인 MustCompile을 사용할 수 있다.
	// Complie은 반환값이 두 개라 상수로 사용할 수 없다.
	rgxMustComplile := regexp.MustCompile("p([a-z]+)ch")
	logger(rgxMustComplile)

	// Func 변형을 사용해서 매칭된 텍스트를 변환시킬 수도 있음.
	in := []byte("a peach")
	out := rgxMustComplile.ReplaceAllFunc(in, bytes.ToUpper)
	logger(string(out))

}
