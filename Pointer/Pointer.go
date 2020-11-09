package main

import (
	"fmt"
)

func main(){
	fmt.Println(`===== 생각해보니 포인터는 굉장히 오랜만이구나 껄껄 =====`)


	tmpA := 0
	tmpB := 1
	tmpAPtr := &tmpA

	fmt.Println(`tmpA : `, tmpA);
	fmt.Println(`tmpB : `, tmpB)
	fmt.Println(`tmpA address : `, tmpAPtr)
	fmt.Println(`tmpB address : `, &tmpB)

	
	tmpB = *tmpAPtr
	fmt.Println(`change ! ! ! `)
	fmt.Println(`tmpA : `, tmpA);
	fmt.Println(`tmpB : `, tmpB )
	fmt.Println(`tmpB address : `, &tmpB )



}