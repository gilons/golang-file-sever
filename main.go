package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	var routes = mux.NewRouter()
	routes.HandleFunc("/photo/save", savePhoto).Methods("POST")
	routes.HandleFunc("/photo/get/{filename}", getPhoto).Methods("GET")
	routes.HandleFunc("/photo/thumbnail/get/{filename}", getPhotoThumbnail).Methods("GET")
	routes.HandleFunc("/sound/save", saveSound).Methods("POST")
	routes.HandleFunc("/sound/get/{filename}", getSound).Methods("GET")
	routes.HandleFunc("/video/save", saveVideo).Methods("POST")
	routes.HandleFunc("/video/get/{filename}", getVideo).Methods("GET")
	routes.HandleFunc("/video/thumbnail/get/{filename}", getVideoThumbnail).Methods("GET")
	routes.HandleFunc("/other/save", saveOther).Methods("POST")
	routes.HandleFunc("/other/get/{filename}", getOther).Methods("GET")

	http.Handle("/", routes)
	log.Fatal(http.ListenAndServe(":8555", nil))

}

func writeThumbnail(fileName string) {
	width := 640
	height := 360
	fileNameArr := strings.Split(fileName, ".")
	outPutFile := fileNameArr[0] + "_thumbnail." + "jpeg"
	cmd := exec.Command("ffmpeg", "-i", fileName, "-vframes", "1", "-an", "-s",
		fmt.Sprintf("%dx%d", width, height), "-ss", "1", outPutFile)
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	err := cmd.Run()
	if err != nil {
		fmt.Println("could not generate frame  Deu to :", err.Error())
	}
}

func getDuration(fileName string) int {
	cmd := exec.Command(`mediainfo`, `--Inform`, `"General;%Duration%"`, fileName)
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	err := cmd.Run()
	if err != nil {
		fmt.Println("could not generate frame  Deu to :", err.Error())
	}
	arr := strings.Split((buffer.String()), "\n")
	for _, element := range arr {
		temp := strings.Split(strings.TrimSpace(element), ":")
		if strings.Contains(temp[0], "Duration") {
			duration := temp[1]
			duration = strings.Join(strings.Split(duration, " "), "")
			durationArr := strings.Split(duration, "min")
			if len(durationArr) > 1 {
				secss := 0
				mins, _ := strconv.ParseInt(durationArr[0], 10, 64)
				secs, _ := strconv.ParseInt(strings.Split(durationArr[1], "s")[0], 10, 64)
				secss = int(mins)*60 + int(secs)
				return secss
			} else {
				durationArr = strings.Split(duration, "s")
				secs, _ := strconv.ParseInt(durationArr[0], 10, 64)
				mms, _ := strconv.ParseInt(strings.Split(durationArr[1], "m")[0], 10, 64)
				return int(secs) + int(mms)/1000
			}
		}
	}
	return 0
}
