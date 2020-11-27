package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mydb "../MySQL"
)

func Middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(`Middleware Test Log : `, r.URL.Path)
		f(w, r)
	}
}

func ReqTest1(w http.ResponseWriter, r *http.Request) {

	resJson, _ := json.Marshal(mydb.DBProcess())

	fmt.Fprintf(w, string(resJson))
}

func main() {
	//라우팅
	http.HandleFunc("/reqeustTest1", Middleware(ReqTest1))

	fmt.Println(`GO LANG WEB SERVER WORKING !`)

	http.ListenAndServe(":8080", nil)
	"git test"
}
