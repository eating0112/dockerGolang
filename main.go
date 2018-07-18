package main

import (
	"fmt"
	"net/http"
)

//type Hello struct{}

//func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "<H1>你好!...</H2>")
//}

func main() {
	fmt.Println("Success")
	//var h Hello
	//http.ListenAndServe("0.0.0.0:8066", h)
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	http.Handle("/handle", http.HandlerFunc(say))
	http.ListenAndServe(":8001", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "<H2> Hello call /hello </H2>")
}

func say(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "<h2> Hello call /handle </h2>")
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "HOME Page !! Can call /hello or /handle")
}
