package main

import (
	"fmt"
	"log"
	"net/http"
)

	func helloHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html>")
		fmt.Fprintf(w, "<head>")
		fmt.Fprintf(w, "</head>")
		fmt.Fprintf(w, "<body>")
		fmt.Fprintf(w, "<h1>Hello</h1>")
		fmt.Fprintf(w, "</body>")
		fmt.Fprintf(w, "</html>")
	}


func main() {

	fmt.Println("Setting handlers... ")
	
	http.HandleFunc("/hello", helloHandler)
	
	err := http.ListenAndServe(":80", nil)
	
	log.Fatal(err)
}