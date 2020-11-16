package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (a ByLength) Len() int           { return len(a) }
func (a ByLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLength) Less(i, j int) bool { return a[i] < a[j] }
func main() {
	fmt.Println(`===== SORT =====`)

	enArr := []string{`a c`, `c`, `b`, `e`, `d`}
	koArr1 := []string{`ㄱ`, `ㄷ`, `ㅁ`, `ㄴ`, `ㄹ`}
	koArr2 := []string{`ㄱㅏ`, `ㄷㅏ`, `ㅁㅏ`, `ㄴㅏ`, `ㄹㅏ`}
	numArr := []int{731, 999, 123, 2346, 1}
	sort.Strings(enArr)
	sort.Strings(koArr1)
	sort.Strings(koArr2)
	sort.Ints(numArr)
	fmt.Println(`EN :`, enArr, `/ KO_1 :`, koArr1, `/ KO_2 :`, koArr2, `/ NUM :`, numArr)

	wordArr := []string{`bb`, `dddd`, `ffffff`, `ggggggg`, `a`, `eeeee`, `ccc`}
	sort.Sort(ByLength(wordArr))
	fmt.Println(wordArr)

}
