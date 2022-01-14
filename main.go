package main

import (
	"fmt"
	"html/template"
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
	type IndexData struct {
		Title string
	}
	path := r.URL.Path
	var fileData string
	templateHome := "./templates/home.html"
	if strings.Index(path, "/static/") != 0 {
		if path == "/" {
			files := append([]string{"./public/index.html"}, templateHome)
			fmt.Println(files)
			t, _ := template.ParseFiles(files...)
			t.Execute(w, "Golang HTTP Server Example")
		} else {
			files := append([]string{fmt.Sprintf("./public%s.html", path)}, templateHome)
			fmt.Println(files)
			t, _ := template.ParseFiles(files...)
			t.Execute(w, nil)
		}
	}
	io.WriteString(w, fileData)
}

func h2(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path != "/static/" {
		http.ServeFile(w, r, "./public"+path)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handleUri)
	http.HandleFunc("/static/", h2)
	log.Fatal(http.ListenAndServe(":3010", nil))
}
