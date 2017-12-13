package main

import (
	"log"
	"os"
	"text/template"
)

func test1() {
	
    //define the data (model)
	student := Student{"Bart Simpson"}
	
	//create a new template with a name
	tmpl := template.New("mytemplate")
	
	// the template content 
	content := "Hello {{.Name}}! \n\n"
	
	//parse some content and generate a template
    tmpl, err := tmpl.Parse(content)
    if err != nil {
		log.Fatal("Parse error : ", err)
		return
	}
    
    // merge the template 'tmpl' with the data
	err1 := tmpl.Execute(os.Stdout, student)
	if err1 != nil {
		log.Fatal("Execute error : ", err1)
		return
	}
}