package main

import (
	"fmt"
	"os"
	"text/template"
)

func testOneHtmlFileBis() {
	
	fmt.Println("Test with a single html template file")
	
    //define the data (model)
	employee := Employee{123, "Bart", "Simpson", 2, true}
	
	tmpl := template.New("my template") // Create a template.
	fmt.Println("Template created, name = '", tmpl.Name(), "'" )
	
	tmpl, err := tmpl.ParseFiles("templates/html/employee1.gohtml")
	if err != nil {
	    panic(err)
	}	
	fmt.Println("Template parsed, name = '", tmpl.Name(), "'" )
	
	fmt.Println("Execute result : " )
	// The template name is the file name
	err = tmpl.ExecuteTemplate(os.Stdout, "employee1.gohtml", employee)
	if err != nil {
	    panic(err)
	}
	
	
//	err = tmpl.Execute(os.Stdout, employee) // Doesn't work
//	if err != nil {
//	    panic(err)
//	}
    
}