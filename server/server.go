package server

import (
	"groupie/server/model"
	"html/template"
	"log"
	"net/http"
)

//create a variable from server package
var Tpl *template.Template
var artists []model.Band

func Start() error {
	artists = handleJSON()
	// content, err := json.Marshal(artists)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = ioutil.WriteFile("artists.json", content, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	tmpl, err := template.ParseFiles("server/template/index.html")
	if err != nil {
		log.Fatal(err)
	}
	//read the template
	// Http. handle(pattern, handler Http.handler) => http.Fileserver(root) return handler => dir : root directory
	// call variable from server package and add value
	Tpl = tmpl
	//handle css from static directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./server/static/"))))
	//get request
	http.HandleFunc("/", GetRequest)
	http.HandleFunc("/artists", GetArtistById)
	//post request

	//open port- listen
	log.Println("Staring server on port...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}
