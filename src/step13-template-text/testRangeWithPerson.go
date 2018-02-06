package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var person = Person{
        Name:   "Foo",
        Emails: []string{"foo@myhost.org", "foo@gmail.com"},
    }

func execute(content string, data interface{} ) {

	fmt.Println("--------------------------------")
	
	//create a new template with a name
	tmpl := template.New("mytemplate")

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
	fmt.Println("--------------------------------")
}

func test1() {
	// the template content 
	content := `Person : {{.Name}}  e-mails :
{{range $e := .Emails}} - email {{$e}} 
{{end}}
`
	execute(content, person)
}

func test2() {
	// the template content 
	content := `Person : {{.Name}}  e-mails :
{{range $i, $e := .Emails}} - email {{index $i}} {{.}} 
{{end}}
`
	execute(content, person)
}

func test3() {
	// the template content 
	content := `Person : {{.Name}}  e-mails :
{{range .Emails}} - email : {{.}} 
{{end}}
`
	execute(content, person)
}

func test4() {
	// the template content 
	content := `{{range .Emails}} - email : {{.}} (name = {{$.Name}})
{{end}}
`
	execute(content, person)
}

func main() {
	
	fmt.Println("Templates tests : ")
	
	test1()
	test2()
	test3()
	test4()
}