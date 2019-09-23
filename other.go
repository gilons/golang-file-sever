package main

import (
	"net/http"
)

const (
	othersFolder = "others/"
)

func getOther(w http.ResponseWriter, r *http.Request) {
	processGet(w, r, othersFolder)
}

func saveOther(w http.ResponseWriter, r *http.Request) {
	processSave(w, r, othersFolder)
}
