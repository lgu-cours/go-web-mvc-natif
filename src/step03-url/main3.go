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

	func urlHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL fields : \n")
		fmt.Fprintf(w, " . URL.Path : %s\n", r.URL.Path)
		fmt.Fprintf(w, " . URL.Scheme : %s\n", r.URL.Scheme)
		fmt.Fprintf(w, " . URL.Host : %s\n", r.URL.Host)
		fmt.Fprintf(w, "URL functions : \n")
		fmt.Fprintf(w, " . URL.Hostname() : %s\n", r.URL.Hostname())
		fmt.Fprintf(w, " . URL.Port() : %s\n", r.URL.Port())
		fmt.Fprintf(w, " . URL.RequestURI() : %s\n", r.URL.RequestURI())
		fmt.Fprintf(w, " . URL.Query() : %s\n", r.URL.Query())
	}


func main() {

	fmt.Println("Setting handlers... ")
	
	http.HandleFunc("/hello", helloHandler)
	
	http.HandleFunc("/url", urlHandler)
	
	err := http.ListenAndServe(":80", nil)
	
	log.Fatal(err)
}