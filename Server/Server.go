package main

import (
	"fmt"
	"log"
	"net/http"
)

func Middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(`Middleware Test Log : `, r.URL.Path)
		f(w, r)
	}
}

func ReqTest1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request : /reqeustTest1")
}

func ReqTest2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request : /reqeustTest2")
}

func main() {
	//라우팅
	http.HandleFunc("/reqeustTest1", Middleware(ReqTest1))
	http.HandleFunc("/reqeustTest2", Middleware(ReqTest2))

	fmt.Println(`GO LANG WEB SERVER WORKING !`)

	http.ListenAndServe(":8080", nil)
}
