package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	
	fmt.Println("Test : ")
	
	data := Person{
        Name:   "Foo",
        Emails: []string{"foo@myhost.org", "foo@gmail.com"},
    }
	
	//create a new template with a name
	tmpl := template.New("mytemplate")
	
	// the template content 
	content := `Person : {{.Name}}  e-mails :
{{range $e := .Emails}} - email {{$e}} 
{{end}}
`
/*
	content := `Person : {{.Name}}  e-mails :
{{range $i, $e := .Emails}} - email {{index $i}} {{.}} 
{{end}}
`
*/

/*
	content := `Person : {{.Name}}  e-mails :
{{range .Emails}} - email : {{.}} 
{{end}}
`
*/

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