package main

import "fmt"

func main() {
	fmt.Println(`===========================================`)
	fmt.Println(`Golang Data Study`)
	fmt.Println(`===========================================`)

	/*
		Array, Slice
	*/
	// 배열 선언
	goArr := [3]int{1, 2, 3}
	// 슬라이스 선언
	goSlice := []int{1, 2, 3}

	fmt.Println(`goArr length :`, len(goArr))
	fmt.Println(`goSlice length :`, len(goSlice))

	// 배열과 슬라이스의 차이?
	// 슬라이스는 동적으로 늘리고 줄일 수 있다.
	// goArr = append(goArr, 4, 5, 6) // 컴파일 에러 발생.
	goSlice = append(goSlice, 4, 5, 6)

	fmt.Println(`goArr length :`, len(goArr))
	fmt.Println(`goSlice length :`, len(goSlice))

	// 간단한 loop
	// golang 에서는 for 로 해결가능한것이 많다! 좋다! 편하네!

	// Three-component loop
	fmt.Println(`===========================================`)
	for idx, arrLen := 0, len(goArr); idx < arrLen; idx++ {

		fmt.Println(`goArr's [`, idx, `] values is `, goArr[idx])
	}

	fmt.Println(`===========================================`)
	for idx, sliceLen := 0, len(goSlice); idx < sliceLen; idx++ {
		fmt.Println(`goSlice's [`, idx, `] values is `, goSlice[idx])
	}

	// for each range loop
	fmt.Println(`===========================================`)
	for idx, value := range goArr {
		fmt.Println(`goArr's [`, idx, `] values is `, value)
	}

	fmt.Println(`===========================================`)
	for idx, value := range goSlice {
		fmt.Println(`goSlice's [`, idx, `] values is `, value)
	}

	/*
		map
	*/

	goMap := make(map[string]string)
	goMap[`title`] = `go lang study`
	goMap[`writer`] = `jay`
	goMap[`content`] = `golang 을 공부중인 go린이의 아무말 블라블라블라`
	goMap[`datetime`] = `2020-11-08 00:29:29 일요일`
	goMap[`delete`] = `이것은 delete 될 예정입니다~`

	fmt.Println(`===========================================`)
	fmt.Println(`before delete map len:`, len(goMap), `(`, goMap, `)`)
	delete(goMap, "delete")
	fmt.Println(`===========================================`)
	fmt.Println(`after delete map len:`, len(goMap), `(`, goMap, `)`)
	fmt.Println(`===========================================`)
	for key, value := range goMap {
		fmt.Println(key, ` -> `, value)
	}

}
