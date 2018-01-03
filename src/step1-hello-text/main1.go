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
	
	http.HandleFunc("/hello", helloHandler)
	
	// If go 1.9+ cannot use port < 1024 
	err := http.ListenAndServe(":80", nil)
	
	log.Fatal(err)
}