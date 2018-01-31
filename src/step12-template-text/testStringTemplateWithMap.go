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
    fmt.Println(" - data is defined in a map ")

	//define the data (model)
	data := map[string]string{"a": "AA",  "b": "BB",  "c": "CC",  "d": "DD"}
	fmt.Println("Value is a map " )
	
	//create a new template with a name
	tmpl := template.New("mytemplate")
	
	// the template content 
	//content := "Test with a single point : Hello {{.}} ! \n\n" // print the entire map
	content := "Test a = {{.a}}, b = {{.b}} \n\n" // print the entire map
	
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