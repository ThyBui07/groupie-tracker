package main

import (
	"groupie/server"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmpl *template.Template

func init() {
	//read the template
	_, err := os.Stat("template/index.html")
	if err == nil {
		tmpl = template.Must(template.ParseFiles("template/index.html"))
	} else {
		tmpl = nil
	}

}
func main() {
	server.Tpl = tmpl
	//handle css from static directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	//get request
	http.HandleFunc("/", server.GetRequest)

	//open port- listen
	log.Fatal(http.ListenAndServe(":8080", nil))

}
