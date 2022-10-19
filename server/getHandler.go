package server

import (
	"fmt"
	"html/template"
	"net/http"
)

//create a variable from server package
var Tpl *template.Template

func GetRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello get")

	Tpl.Execute(w, "hello")
}
