package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type MyKeyBoard struct {
	Brand  string
	Switch string
	keyCnt int
}

func ErrCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	ErrCheck(err)
	fmt.Println(`=============== READ FILE RESULT =================`)
	fmt.Println(string(data))
	fmt.Println(`==================================================`)
}

func WriteFile(filePath string, content string) {
	err := ioutil.WriteFile(filePath, []byte(content), os.FileMode(644))
	ErrCheck(err)
}

func main() {

	filePath := `/Users/jaylee/go/ToyProject/FileIO/wrtie_file_test.json`
	// Go 데이타
	myKeyBoardList := []MyKeyBoard{
		MyKeyBoard{`KEYCHRON`, `RED`, 104},
		MyKeyBoard{`VAMILO`, `SLIENT RED`, 108},
		MyKeyBoard{`LEOPOLD_FC_980C`, `TOPRE`, 97},
		MyKeyBoard{`LOGITECH`, `Membrain`, 104},
		MyKeyBoard{`HANSUNG`, `Noppoo`, 108},
	}

	// JSON 인코딩
	jsonBytes, err := json.Marshal(myKeyBoardList)
	ErrCheck(err)

	// JSON 바이트를 문자열로 변경
	jsonString := string(jsonBytes)

	WriteFile(filePath, jsonString)
	ReadFile(filePath)

}
