package main

import (
	"os"
	"text/template"
)

func testOneFile() {
	
    //define the data (model)
	student := Student{"Bart Simpson"}
	
	tmpl, err := template.ParseFiles("templates/test2.tpl")
	if err != nil {
	    panic(err)
	}
	// The template name is the file name
	err = tmpl.ExecuteTemplate(os.Stdout, "test2.tpl", student)
	if err != nil {
	    panic(err)
	}
    
}