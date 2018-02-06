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
	data := "Bart Simpson"
	fmt.Println("Value is : " + data )
	
	//create a new template with a name
	tmpl := template.New("mytemplate")
	
	// the template content 
	content := "First test : \n" +
			 "Hello {{.}} ! \n\n" +
			 // "Add 1 + 2  {{ 1 + 2}} ! \n\n" +
			 "{{ $a := 1 }} a = {{ $a}} \n\n" 
			 
	
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