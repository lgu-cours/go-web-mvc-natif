package main

import (
	"fmt"
	"log"
	"net/http"
)

func genHtmlPage(w http.ResponseWriter, content string) {
	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<head>")
	fmt.Fprintf(w, "</head>")
	fmt.Fprintf(w, "<body>")
	fmt.Fprintf(w, content)
	fmt.Fprintf(w, "</body>")
	fmt.Fprintf(w, "</html>")
}
	
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	
	content := `
	<h1>Welcome</h1>
	<ul>
	<li><a href="param"        >/param</a> </li>
	<li><a href="param?key=aa" >/param?key=aa</a> </li>
	<li><a href="param?key="   >/param?key=</a> </li>
	<li><a href="param?key=aa&key=bb" >/param?key=aa&key=bb</a> </li>
	</ul>
	`
	genHtmlPage(w, content)
}
	
func getParameter(request *http.Request, name string) string {
	
	// r.URL.Query() returns a 'Values' type
	// which is simply a map[string][]string of the QueryString parameters.
	queryValues := request.URL.Query()
	
    // Query()["key"] will return an array of items, 
    // we only want a single item => use the first one
	values, ok := queryValues[name]
    if ok && len(values) > 0 {
	    return values[0]
    }
	return ""
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	
	keyValue := getParameter(r, "key")
	
    log.Println("Url Param 'key' = '" + keyValue + "'")
    if ( keyValue != "" ) {
	    s := fmt.Sprintf("<h2>'key' value is %s</h2>", keyValue)
	    genHtmlPage(w, s)
    } else {
	    genHtmlPage(w, "<h2>No 'key' parameter or void value")
    }
}

func main() {

	fmt.Println("Setting handlers... ")
	
	// By default the welcome page
	http.HandleFunc("/", welcomeHandler)			

	http.HandleFunc("/param", paramHandler)

	err := http.ListenAndServe(":80", nil)
	
	log.Fatal(err)
}