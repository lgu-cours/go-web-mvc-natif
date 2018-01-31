package main

import (
	"fmt"
	"log"
	"net/http"
)

const webDir = "www"

func serveStaticFile(w http.ResponseWriter, r *http.Request) {	
	s := webDir + r.URL.Path
	log.Print("Static file requested. URL path '" + r.URL.Path + "' --> serve file '" + s + "'" )
	http.ServeFile(w, r, s)
}

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
	
	// By default serve static files
	http.HandleFunc("/", serveStaticFile)			

	//http.HandleFunc("/hello", helloHandler) // "/hello" strict
	http.HandleFunc("/hello/", helloHandler) // "/hello/*"
	
	err := http.ListenAndServe(":80", nil)
	
	log.Fatal(err)
}