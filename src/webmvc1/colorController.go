package main

import (
	"fmt"
	"log"
	"net/http"
)

func colorController(w http.ResponseWriter, r *http.Request) {
	log.Print("colorController - URL path '" + r.URL.Path )

    // func (u *URL) Query() Values {
    queryValues := r.URL.Query()
    name := queryValues.Get("name")
    color := queryValues.Get("color")
    
//    fmt.Fprintf(w, "name = %s\n", name)
//    fmt.Fprintf(w, "color = %s\n", color)
    
    fmt.Fprint(w, "<!DOCTYPE html>")
    fmt.Fprint(w, "<html>")
    
    fmt.Fprint(w, "<body>")
    fmt.Fprintf(w, "<h2>Hello %s ! </h2>\n", name)
    fmt.Fprintf(w, "<h2>You prefered color is %s </h2>\n", color)
    
    fmt.Fprintf(w, "<a href=\"color.html\" >Play again...</a>");
    fmt.Fprint(w, "</body>")

    fmt.Fprint(w, "</html>")

}
