package server

import (
	"fmt"
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
	// res := model.Locations{}
	// artistsData := getAPI("https://groupietrackers.herokuapp.com/api/locations")
	// json.Unmarshal([]byte(artistsData), &res)
	fmt.Println("hello get")
	//json.NewEncoder(w).Encode(peter)
	Tpl.Execute(w, artists)
}
