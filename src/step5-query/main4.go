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
	
func paramHandler(w http.ResponseWriter, r *http.Request) {
	
	// r.URL.Query() returns a Values type, 
	// which is simply a map[string][]string of the QueryString parameters.
	keys, ok := r.URL.Query()["key"]
    
    if !ok || len(keys) < 1 {
        log.Println("Url Param 'key' is missing")
        fmt.Fprintf(w, "<h2>Param 'key' is missing</h2>")
        return
    }

    // Query()["key"] will return an array of items, 
    // we only want the single item.
    log.Printf(" len(keys) = %d", len(keys) )
    keyValue := keys[0]

    log.Println("Url Param 'key' is " + keyValue )
    fmt.Fprintf(w, "<h2>'key' value is %s</h2>", keyValue)

}


func main() {

	fmt.Println("Setting handlers... ")
	
	// By default serve static files
	http.HandleFunc("/", serveStaticFile)			

	//http.HandleFunc("/hello", helloHandler) // "/hello" strict
	http.HandleFunc("/hello/", helloHandler) // "/hello/*"
	
	http.HandleFunc("/param", paramHandler)

	err := http.ListenAndServe(":80", nil)
	
	log.Fatal(err)
}