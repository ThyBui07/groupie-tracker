package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
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
	// res := model.Locations{}
	// artistsData := getAPI("https://groupietrackers.herokuapp.com/api/locations")
	// json.Unmarshal([]byte(artistsData), &res)
	fmt.Println("hello get")
	//json.NewEncoder(w).Encode(peter)
	Tpl.Execute(w, artists)
}

func GetArtistById(w http.ResponseWriter, r *http.Request) {
	Tpl, err := template.ParseFiles("server/template/band.html")
	if err != nil {
		log.Fatal(err)
	}
	id := r.URL.Query().Get("id")
	num, e := strconv.Atoi(id)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("id =>", id)
	res := artists[(num - 1)]
	Tpl.Execute(w, res)
}
