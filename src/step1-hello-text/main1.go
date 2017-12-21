package main

import (
	"fmt"
	"log"
	"net/http"
)

	func helloHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello !")
	}


func main() {

	fmt.Println("Setting handlers... ")
	
	http.HandleFun

	http.HandleFunc("/hello", helloHandler)
	
	err := http.ListenAndServe(":80", nil)
	
	log.Fatal(err)
}