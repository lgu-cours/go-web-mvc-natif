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
    fmt.Println(" - data is defined in a simple string ")

	//define the data (model)
	value := "Bart Simpson"
	fmt.Println("Value is : " + value )
	
	//create a new template with a name
	tmpl := template.New("mytemplate")
	
	// the template content 
	content := "Test with a single point : Hello {{.}} ! \n\n"
	
	//parse some content and generate a template
    tmpl, err := tmpl.Parse(content)
    if err != nil {
		log.Fatal("Parse error : ", err)
		return
	}
    
    // merge the template 'tmpl' with the data
	err1 := tmpl.Execute(os.Stdout, value)
	if err1 != nil {
		log.Fatal("Execute error : ", err1)
		return
	}
}