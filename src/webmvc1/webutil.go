package main

import (
	"fmt"
	"net/http"
	"strings"
)

func printFormValues(w http.ResponseWriter, r *http.Request) {
 
    r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
 	
	fmt.Fprintf(w, "Form values : \n")
    for k, v := range r.Form {
        fmt.Fprintf(w, " '%s' --> '%s' \n", k, strings.Join(v, ",") )
    }
}
