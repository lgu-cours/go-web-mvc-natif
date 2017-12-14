package main

import (
	"fmt"
	"log"
	"net/http"
)

func form1Controller(w http.ResponseWriter, r *http.Request) {
	log.Print("form1Controller - URL path '" + r.URL.Path )

	fmt.Fprintf(w, "form1 controller response \n")

    // func (u *URL) Query() Values {
    queryValues := r.URL.Query()
    firstName := queryValues.Get("firstName")
    lastName := queryValues.Get("lastName")
    x := queryValues.Get("x")
    
    fmt.Fprintf(w, "firstName = %s\n", firstName)
    fmt.Fprintf(w, "lastName = %s\n", lastName)
    fmt.Fprintf(w, "x = %s\n", x)
    
    // Get("foo") : 2 values => get only the first one
    fmt.Fprintf(w, "foo (with Get function) = %s\n", queryValues.Get("foo"))
    
//    // func (u *URL) Query() Values {
//    values := r.URL.Query()
    
    // type Values map[string][]string
    fooValues := queryValues["foo"]
    fmt.Fprintf(w, "foofooValues (with x['foo'] = %s\n", fooValues)
    
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
