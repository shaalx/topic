package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/auth/qq/signin", signinHandler)
	http.ListenAndServe(":4000", nil)
	fmt.Println("start...")
}

func signinHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("signinHandler.."))
}
