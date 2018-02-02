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
	<form action="/param" method="POST">
	<table>
	<tr><td>First name : </td> <td> <input type="text" name="firstname" placeholder="first name"> </td> </tr>
	<tr><td>Last name :  </td> <td> <input type="text" name="lastname" placeholder="last name"> </td> </tr>
	</table>
	<button type="submit">Submit</button>
	</form>	
	`
	genHtmlPage(w, content)
}
	
func paramHandler(w http.ResponseWriter, r *http.Request) {
	
    r.ParseForm() // Parse url parameters passed, then parse the POST body (request body)
    firstname := r.Form.Get("firstname")
    lastname  := r.Form.Get("lastname")

    log.Println("Param 'firstname' = '" + firstname + "'")
    log.Println("Param 'lastname'  = '" + lastname  + "'")

    s := fmt.Sprintf("<h2>'firstname' value is '%s'</h2> <h2>'lastname' value is '%s'</h2>", firstname, lastname)
    genHtmlPage(w, s)
}

func main() {

	fmt.Println("Setting handlers... ")
	
	http.HandleFunc("/", welcomeHandler)			

	http.HandleFunc("/param", paramHandler)

	err := http.ListenAndServe(":80", nil)
	
	log.Fatal(err)
}