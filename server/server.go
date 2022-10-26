package server

import (
	"encoding/json"
	"fmt"
	"groupie/server/model"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

//create a variable from server package
var Tpl *template.Template
var artists []model.Band

// func getData() {
// 	res := handleJSON()
// 	fmt.Printf("this is res type : %T", res)
// 	content, err := json.Marshal(res)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = ioutil.WriteFile("artists.json", content, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
func Start() error {
	artists = handleJSON()
	content, err := json.Marshal(artists)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("artists.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.ParseFiles("server/template/index.html")
	if err != nil {
		return err
	}
	//read the template
	// Http. handle(pattern, handler Http.handler) => http.Fileserver(root) return handler => dir : root directory
	// call variable from server package and add value
	Tpl = tmpl
	//handle css from static directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/server/static/"))))
	//get request
	http.HandleFunc("/", GetRequest)
	//post request

	//open port- listen
	log.Println("Staring server on port...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}
