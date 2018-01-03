package main

import (
	"fmt"
	"log"
	"net/http"
	
	"webmvc3/controllers"
)

const webDir = "www"

func serveStaticFile(w http.ResponseWriter, r *http.Request) {		
	s := webDir + r.URL.Path
	log.Print("Static file requested. URL path '" + r.URL.Path + "' --> server file '" + s + "'" )
	http.ServeFile(w, r, s)
}

func main() {
	printDir()

	fmt.Println("Setting handlers... ")
	
	// By default serve static files
	http.HandleFunc("/", serveStaticFile)			

	// Specific Paths with specific controllers 
	http.HandleFunc("/color", controllers.ColorController)

	log.Fatal(http.ListenAndServe(":80", nil))

}