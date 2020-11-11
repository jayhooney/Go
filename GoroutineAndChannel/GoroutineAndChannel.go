	
package main
import "fmt"

func GoFunc(writer chan string) {
	// 채널에 값 전달
	writer <- `jay`
	writer <- `hoon`
}

func main() {

	// 채널 생성
	writer := make(chan string, 2)


	go GoFunc(writer)

	// 채널 값 받아오기
	// myName := <- writer
	fmt.Println(`my name is`, <- writer)
	fmt.Println(`my name is`, <- writer)

}