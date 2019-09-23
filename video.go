package main

import (
	"net/http"
)

const (
	videoFolder = "videos/"
)

func saveVideo(w http.ResponseWriter, r *http.Request) {
	processSave(w, r, videoFolder)
}

func getVideo(w http.ResponseWriter, r *http.Request) {
	processSave(w, r, videoFolder)
}
