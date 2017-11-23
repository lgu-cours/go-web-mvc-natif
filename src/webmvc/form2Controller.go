package main

import (
	"fmt"
	"net/http"
)

func form2Controller(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "form2 controller response \n")
    r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Fprintf(w, "Form : %s \n", r.Form ) // print information on server side.
    fmt.Fprintf(w, "URL.Path : %s \n", r.URL.Path)
    fmt.Fprintf(w, "URL.Scheme :%s \n", r.URL.Scheme )
    
    // r.Form contient les valeurs POST et les valeurs de QUERY String de l'URL 
    printFormValues(w,r);
    
    fmt.Fprintf(w, "firstName = %s \n", r.Form.Get("firstName") )
    fmt.Fprintf(w, "lastName = %s \n", r.Form.Get("lastName") )

    fmt.Fprintf(w, "foo = %s \n", r.Form.Get("foo") )
    
}