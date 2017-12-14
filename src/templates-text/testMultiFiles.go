package main

import (
	"fmt"
	"os"
	"text/template"
)

func testMultiFiles() {
	
	fmt.Println("Test with multiple templates files parsed once at the beginning")

    //define the data (model)
	student := Student{"Bart Simpson"}
	
	tmpl, err := template.ParseFiles("templates/text/test2.gotxt", "templates/text/test3.gotxt")
	if err != nil {
	    panic(err)
	}
	// The template name is the file name
	err = tmpl.ExecuteTemplate(os.Stdout, "test2.gotxt", student)
	if err != nil {
	    panic(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "test3.gotxt", student)
	if err != nil {
	    panic(err)
	}
    
}