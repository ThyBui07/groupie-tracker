package server

import (
	"errors"
	"fmt"
	"groupie/server/model"
	"html/template"
	"log"
	"net/http"
)

func GetRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorNotFound(w)
		return
	}
	if Tpl == nil {
		errorInternalServer(w)
		return
	}
	Tpl.Execute(w, artists)
}

func GetArtistById(w http.ResponseWriter, r *http.Request) {
	Tpl, err := template.ParseFiles("server/template/band.html")
	if err != nil {
		log.Fatal(err)
	}
	query := r.URL.Query()
	fmt.Println(query)
	name := r.URL.Query().Get("name")
	var res model.Band
	if name == "" {
		errorBadRequest(w, errors.New("invalid query"))
		return
	} else {
		for i, v := range artists {
			if v.Name == name {
				res = artists[i]
				break
			}
		}
	}
	if res.Id == 0 {
		errorNotFound(w)
		return
	}
	Tpl.Execute(w, res)
}
