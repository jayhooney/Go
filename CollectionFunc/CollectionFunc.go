package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for idx, v := range vs {
		if v == t {
			return idx
		}
	}

	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}

	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}

	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))

	for idx, v := range vs {
		vsm[idx] = f(v)
	}

	return vsm
}

func main() {
	fmt.Println(`===== COLLECTION FUNC =====`)

	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	// 위의 예시 모두 익명함수를 사용했지만, 올바른 타입을 가진 명명된 함수를 사용할 수도 있습니다.

	fmt.Println(Map(strs, strings.ToUpper))
}
