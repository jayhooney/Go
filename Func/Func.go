package main

import (
	"fmt"
)

//Variadic functions

func Variadic(intVals ...int) int {
	fmt.Println(`1. 가변 함수`)

	fmt.Println(`파라미터 길이 : `, len(intVals))
	fmt.Println(`파라미터 : `, intVals)

	sum := 0

	for idx, len := 0, len(intVals); idx < len; idx++ {
		fmt.Println(`[`, idx, `] 번째 파라미터 =>`, intVals[idx])
		sum = sum + intVals[idx]
	}

	return sum
}

func MultipleReturnValues(params ...int) (int, string) {
	fmt.Println(`2. 다중 반환 함수 `)

	fmt.Println(`파라미터 길이 : `, len(params))
	fmt.Println(`파라미터 : `, params)

	sum := 0

	for idx, len := 0, len(params); idx < len; idx++ {
		fmt.Println(`[`, idx, `] 번째 파라미터 =>`, params[idx])
		sum = sum + params[idx]
	}

	msg := `총 합 : `

	return sum, msg

}

func main() {
	fmt.Println(`func study`)

	params := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	varidicResult := Variadic(params...)
	fmt.Println(`총 합 : `, varidicResult)

	multipleResult, msg := MultipleReturnValues(params...)
	fmt.Println(msg, multipleResult)

}
