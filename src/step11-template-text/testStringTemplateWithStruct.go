package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	
	fmt.Println("Test : ")
	fmt.Println(" - template is defined in a string")
    fmt.Println(" - data is defined in a structure ")
    //define the data (model)
	student := Student{"Bart", "Simpson", "abc" }
	
	//create a new template with a name
	tmpl := template.New("mytemplate")
	
	// the template content 
	//content := "Hello {{.FirstName}} {{.LastName}} ! \n Secret code = {{.secretCode}} \n Foo = {{.Foo}} \n"
	content := "Hello {{.FirstName}} {{.LastName}} ! \n Foo = {{.Foo}} \n"
	
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