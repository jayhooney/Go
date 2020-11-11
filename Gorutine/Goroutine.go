package main

import "fmt"

func funcA(from string){
	for idx,len:=0,10; idx < len; idx ++ {
		fmt.Println(`FROM : `, from, `=> `, idx)
	}
}

func main() {

	go funcA(`hoon`)
	funcA(`jay`)

	

	input := ``
	fmt.Scanln(&input)
}