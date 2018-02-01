package main

import (
	"fmt"
	"os"
	"html/template"
)

func testOneHtmlFile() {
	
	fmt.Println("Test with a single html template file")
	
    //define the data (model)
	employee := Employee{123, "Bart", "Simpson", 2, true}
	
	tmpl, err := template.ParseFiles("templates/html/employee1.gohtml")
	if err != nil {
	    panic(err)
	}
	fmt.Println("Template parsed, name = '", tmpl.Name(), "'" )
	fmt.Println("DefinedTemplates() = '", tmpl.DefinedTemplates(), "'" )
	
	fmt.Println("==========" )
	fmt.Println("Execute with 'tmpl.ExecuteTemplate' : " )
	// The template name is the file name
	err = tmpl.ExecuteTemplate(os.Stdout, "employee1.gohtml", employee)
	if err != nil {
	    panic(err)
	}

	fmt.Println("==========" )
	fmt.Println("Execute with 'tmpl.Execute' : " )
	err = tmpl.Execute(os.Stdout, employee)
	if err != nil {
	    panic(err)
	}
    
}