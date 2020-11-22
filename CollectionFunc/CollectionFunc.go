package main

import (
	"fmt"
	"strings"
)

// 문자열 t의 첫 위치를 반환하고 존재하지 않으면 -1를 반환.
func Index(vs []string, t string) int {
	for idx, v := range vs {
		if v == t {
			return idx
		}
	}

	return -1
}

// 문자열 t 가 slice 에 존재하면 true를 반환.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 문자열중 하나가 조건 f를 만족하면 true를 반환.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}

	return false
}

// slice의 문자열 모두가 조건 f를 만족하면 true를 반환.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// slice에서 조건 f를 만족하는 모든 문자열을 포함하는 새 slice를 반환.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}

	return vsf
}

// 기존 slice에 있는 각 문자열에 함수 f 를 적용한 결과값들을 포함해 새 slice를 반환.
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

	fmt.Println(Map(strs, strings.ToUpper))
}
