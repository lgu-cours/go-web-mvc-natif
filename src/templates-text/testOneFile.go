package main

import (
	"os"
	"text/template"
)

func testOneFile() {
	
    //define the data (model)
	student := Student{"Bart Simpson"}
	
	tmpl, err := template.ParseFiles("templates/text/test2.gotxt")
	if err != nil {
	    panic(err)
	}
	// The template name is the file name
	err = tmpl.ExecuteTemplate(os.Stdout, "test2.gotxt", student)
	if err != nil {
	    panic(err)
	}
    
   	// tmpl = template.New("mytemplate") // inutile !!
	tmpl.ParseFiles("templates/text/test2.gotxt")
	//err = tmpl.Execute(os.Stdout, student) // template: mytemplate: "mytemplate" is an incomplete or empty template
	//err = tmpl.ExecuteTemplate(os.Stdout, "mytemplate", student) // no template "mytemplate" associated with template "mytemplate"
	err = tmpl.ExecuteTemplate(os.Stdout, "test2.gotxt", student) // OK
	if err != nil {
	    panic(err)
	}

}