package main

import (
	"fmt"
	Models "Go/StructAndInterface/Models"
)



func main () {
	fmt.Println(`===== 구조체 / 인터페이스 =====`)
	myKeyboard := Models.KeyboardInfo{ Brand:`keychron`,Name:`K1 version 4`,SwitchType:`red`, SwitchCnt:104 }

	text := Models.KeyboardInfoText.Fulltxt(myKeyboard)
	

	fmt.Println(text)

}