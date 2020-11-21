package main

import (
	"fmt"
	"os"
)

func CreateFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func WriteFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func CloseFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func main() {
	fmt.Println(`===== DEFER =====`)
	// defer는 반드시 수행해야하는 로직이 있을 때 사용.
	// try - catch - finally 에서 finally 와 같은 역할을 함.

	filePtr := CreateFile("./GO_LANG/DEFER.txt")
	defer CloseFile(filePtr)
	WriteFile(filePtr)

}
