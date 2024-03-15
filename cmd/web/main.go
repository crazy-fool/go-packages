package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "hello world")
	if err != nil {
		return
	}
}

type Fu func() string

func test() string {
	return "123"
}

func main() {
	//t := test
	//fu := Fu(t)
	//fmt.Printf("%T\r\n", t)
	//fmt.Printf("%T\r\n", fu)
	//fmt.Println(fu)
	//fmt.Println(fu())
	//return
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
