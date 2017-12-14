package main

import (
	"net/http"	
	"html/template"
)

func forward(w http.ResponseWriter, filePath string, data interface{} ) {
	
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
	    panic(err)
	}
	
	err = tmpl.Execute(w, data)
	if err != nil {
	    panic(err)
	}
	
	
}

