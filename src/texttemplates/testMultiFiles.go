package main

import (
	"os"
	"text/template"
)

func testMultiFiles() {
	
    //define the data (model)
	student := Student{"Bart Simpson"}
	
	tmpl, err := template.ParseFiles("templates/test2.tpl", "templates/test3.tpl")
	if err != nil {
	    panic(err)
	}
	// The template name is the file name
	err = tmpl.ExecuteTemplate(os.Stdout, "test2.tpl", student)
	if err != nil {
	    panic(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "test3.tpl", student)
	if err != nil {
	    panic(err)
	}
    
}