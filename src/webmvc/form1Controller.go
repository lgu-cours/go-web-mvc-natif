package main

import (
	"fmt"
	"log"
	"net/http"
)

func form1Controller(w http.ResponseWriter, r *http.Request) {
	log.Print("form1Controller - URL path '" + r.URL.Path )

	fmt.Fprintf(w, "form1 controller response \n")
    queryValues := r.URL.Query()
    fmt.Fprintf(w, "firstName = %s\n", queryValues.Get("firstName"))
    fmt.Fprintf(w, "lastName = %s\n", queryValues.Get("lastName"))
    
    // Get("foo") : 2 values => get only the first one
    fmt.Fprintf(w, "foo = %s\n", queryValues.Get("foo"))
    
    // func (u *URL) Query() Values {
    values := r.URL.Query()
    
    // type Values map[string][]string
    fooValues := values["foo"]
    
//    fooValues, ok := r.URL.Query()["foo"]
//    if !ok || len(fooValues) < 1 {
//        fmt.Fprintf(w, "Url Param 'foo' is missing")
//        return
//    }
    if len(fooValues) < 1 {
        fmt.Fprintf(w, "Url Param 'foo' is missing")
        return
    } else {
    	n := 0 
	    for i, v := range fooValues {
	    	n++
	        fmt.Fprintf(w, " %d (%d) --> '%s' \n", n, i, v )
	    }
    }
}
