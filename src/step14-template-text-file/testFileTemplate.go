package main

import (
	"os"
	"text/template"
)

func testOneFile() {
	
    //define the data (model)
	student := Student{"Bart Simpson", 19}
	
	tmpl, err := template.ParseFiles("templates/step14/test14.gotxt", "templates/step14/test14bis.gotxt")
	if err != nil {
	    panic(err)
	}
	// The template name must be one the files parsed
	err = tmpl.ExecuteTemplate(os.Stdout, "test14.gotxt", student)
	if err != nil {
	    panic(err)
	}
    
	err = tmpl.ExecuteTemplate(os.Stdout, "test14bis.gotxt", student) // OK
	if err != nil {
	    panic(err)
	}

}