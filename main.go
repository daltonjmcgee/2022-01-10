package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "File not found", err
	} else {
		return string(bytes), err
	}
}

func handleUri(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var fileData string
	if strings.Index(path, "/static/") != 0 {
		if path == "/" {
			data, _ := loadFile("./public/index.html")
			fileData = data
		} else {
			data, _ := loadFile(fmt.Sprintf("./public%s.html", path))
			fileData = data
		}
	}
	io.WriteString(w, fileData)
}

func main() {
	http.HandleFunc("/", handleUri)

	fs := http.FileServer(http.Dir("./public/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":3010", nil))
}
