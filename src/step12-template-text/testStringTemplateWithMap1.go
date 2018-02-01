package main

import (
	"log"
	"os"
	"text/template"
)

func testMap1() {
	
	//define the data (model)
	data := map[string]string{"v1": "One",  "v2": "Two",  "v3": "Three",  "v4": "Four"}
	
	//create a new template with a name
	tmpl := template.New("mytemplate")
	
	// the template content 
	//content := "Test with a single point : Hello {{.}} ! \n\n" // print the entire map
	content := "Some map values : v1 = '{{.v1}}' / v2 = '{{.v2}}' \n"
		
	// the template content 
	//content := "Test with a single point : Hello {{.}} ! \n\n" // print the entire map
	//content := "Test a = {{.a}}, b = {{.b}} \n\n" // print the entire map
	
	//parse some content and generate a template
    tmpl, err := tmpl.Parse(content)
    if err != nil {
		log.Fatal("Parse error : ", err)
		return
	}
    
    // merge the template 'tmpl' with the data
	err1 := tmpl.Execute(os.Stdout, data)
	if err1 != nil {
		log.Fatal("Execute error : ", err1)
		return
	}
}