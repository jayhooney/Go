package main

import (
	"fmt"
	Models "Go/Struct/Models"
)



func main () {
	fmt.Println(`===== 구조체 =====`)
	myKeyboard := Models.KeyboardInfo{ Brand:"keychron",Name:`K1 version 4`,SwitchType:"red", SwitchCnt:104 }

	fmt.Println(`Brand :`, myKeyboard.Brand,`/ Name :`, myKeyboard.Name , `/ SwitchType :`, myKeyboard.SwitchType, `/ SwitchCnt :`, myKeyboard.SwitchCnt)

}