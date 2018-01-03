package controllers

import (
	"log"
	"net/http"
	
	"webmvc3/util"
)

func ColorController(w http.ResponseWriter, r *http.Request) {
	log.Print("colorController - URL path '" + r.URL.Path )

	switch r.Method {
	case "GET":
	    doGet(w,r)
	case "POST":
	    doPost(w,r)
	default:
	    util.ErrorPage(w, "Method "+r.Method+ " is not supported");
	}
}

func doGet(w http.ResponseWriter, r *http.Request) {
	// get parameters to init form data
    r.ParseForm() // Parse url parameters passed, then parse the response packet for the POST body (request body)
    name := r.Form.Get("name")
	// init form data containing only the name 
	data := FormData{name, ""}
	// forward to initial page
	util.Forward(w, "templates/html/color1.gohtml", data)
}

func doPost(w http.ResponseWriter, r *http.Request) {
	// get parameters to init form data
    r.ParseForm() // Parse url parameters passed, then parse the response packet for the POST body (request body)
    name := r.Form.Get("name")
    color  := r.Form.Get("color")
	data := FormData{name, color}
	
	// forward to result page
	util.Forward(w, "templates/html/color2.gohtml", data)
}
