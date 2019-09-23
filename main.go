package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var routes = mux.NewRouter()
	routes.HandleFunc("/photo/save", savePhoto).Methods("POST")
	routes.HandleFunc("/photo/get/{filename}", getPhoto).Methods("GET")
	routes.HandleFunc("/sound/save", saveSound).Methods("POST")
	routes.HandleFunc("/sound/get/{filename}", getSound).Methods("GET")
	routes.HandleFunc("/video/save", getVideo).Methods("POST")
	routes.HandleFunc("/video/get/{filename}", getVideo).Methods("GET")
	routes.HandleFunc("/others/save", saveOther).Methods("POST")
	routes.HandleFunc("/other/get/{filename}", getOther).Methods("GET")

	http.Handle("/", routes)
	log.Fatal(http.ListenAndServe(":8555", nil))

}
