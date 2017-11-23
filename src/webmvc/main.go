package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const webDir = "www"

func serveStaticFile(w http.ResponseWriter, r *http.Request) {		
	s := webDir + r.URL.Path
	log.Print("Static file requested. URL path '" + r.URL.Path + "' --> server file '" + s + "'" )
	http.ServeFile(w, r, s)
}

func hiController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}

func form1Controller(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "form1 controller response \n")
    queryValues := r.URL.Query()
    fmt.Fprintf(w, "firstName = %s\n", queryValues.Get("firstName"))
    fmt.Fprintf(w, "lastName = %s\n", queryValues.Get("lastName"))
}

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

func main() {
	printDir()

	fmt.Println("Setting handlers... ")
	
	// By default serve static files
	http.HandleFunc("/", serveStaticFile)			

	// Specific Paths 
	http.HandleFunc("/hi", hiController)
	http.HandleFunc("/form1", form1Controller)
	http.HandleFunc("/form2", form2Controller)

	log.Fatal(http.ListenAndServe(":80", nil))

}