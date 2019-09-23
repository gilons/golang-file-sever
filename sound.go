package main

import (
	"net/http"
)

const (
	soundFolder = "sounds/"
)

func saveSound(w http.ResponseWriter, r *http.Request) {
	processSave(w, r, soundFolder)
}

func getSound(w http.ResponseWriter, r *http.Request) {
	processGet(w, r, soundFolder)
}
