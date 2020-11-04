package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	text := `Hi Go lang. this tis file IO test`

	err := ioutil.WriteFile("/Users/jaylee/go/ToyProject/FileIO/wrtie_file_test.txt", []byte(text), os.FileMode(644))

	errCheck(err)

	data, err := ioutil.ReadFile("/Users/jaylee/go/ToyProject/FileIO/wrtie_file_test.txt")
	errCheck(err)

	fmt.Println(`================================`)
	fmt.Println(string(data))
	fmt.Println(`================================`)
}
