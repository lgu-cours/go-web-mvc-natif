package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"os"
	"path/filepath"	
)

func printFormValues(w http.ResponseWriter, r *http.Request) {
 
    r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
 	
	fmt.Fprintf(w, "Form values : \n")
    for k, v := range r.Form {
        fmt.Fprintf(w, " '%s' --> '%s' \n", k, strings.Join(v, ",") )
    }
}

func errorPage(w http.ResponseWriter, message string) {
    fmt.Fprint(w, "<!DOCTYPE html>")
    fmt.Fprint(w, "<html>")
    
    fmt.Fprint(w, "<body>")
    fmt.Fprintf(w, "<h1>ERROR</h1>\n")
    fmt.Fprintf(w, "<h2>%s</h2>\n", message)
    fmt.Fprint(w, "</body>")

    fmt.Fprint(w, "</html>")
}

func printDir() {
	// Current dir
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
    fmt.Println("bin dir : " + dir)
    
    dir,err = filepath.Abs("./")
    if err != nil {
            log.Fatal(err)
    }
    fmt.Println("file path : " + dir)
}