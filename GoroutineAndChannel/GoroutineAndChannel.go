package main

import (
	"fmt"
)

func GoroutineFuncA(from chan string){
	from <- `jay`
	for idx,len:=0,10; idx < len; idx ++ {
		fmt.Println(`<GoroutineFuncA> FROM : `, <- from, `=> `, idx)
	}
}

func GoroutineFuncB(from chan string){
	from <- `hoon`
	for idx,len:=0,10; idx < len; idx ++ {
		fmt.Println(`<GoroutineFuncB> FROM : `, <- from, `=> `, idx)
	}
}

func main() {

	fmt.Println(` ===== GO ROUTINE AND CHANNEL =====`)


	fromChannel := make(chan string)
	
	go GoroutineFuncB(fromChannel)
	
	go GoroutineFuncA(fromChannel)

	fmt.Println(<- fromChannel)

	

	fmt.Println(`END`)

	

}