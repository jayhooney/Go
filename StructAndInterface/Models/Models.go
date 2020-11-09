package models

import (
	"strconv"
)

type KeyboardInfoText interface {
	Fulltxt() string
}

type KeyboardInfo struct {
	Brand string
	Name string
	SwitchType string
	SwitchCnt int
}

func (keyboardInfo KeyboardInfo) Fulltxt() string {
	return `Brand :`+ keyboardInfo.Brand + ` Name : `+ keyboardInfo.Name+` SwitchType :`+ keyboardInfo.SwitchType+` SwitchCnt :`+ strconv.Itoa( keyboardInfo.SwitchCnt)
}


