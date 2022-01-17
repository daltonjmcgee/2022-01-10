package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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

func handleDynamic(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path != "/static/" {
		http.ServeFile(w, r, "./public"+path)
	} else {
		http.NotFound(w, r)
	}
}

func handle404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	file, err := loadFile("./public/404.html")
	if err != nil {
		fmt.Fprintf(w, "Some dipshit deleted the default 404 and didn't replace it. At any rate, your page wasn't found.")
	} else {
		fmt.Fprintf(w, file)
	}
}

func handleUri(w http.ResponseWriter, r *http.Request) {
	templateHome := "./templates/home.html"
	type IndexData struct {
		Title string
	}
	path := r.URL.Path

	// Checking for pattern used for dynamic pages and return 404 if found.
	// We don't want anyone grabbing that un-rendered page.
	matched, _ := regexp.Match(`\[\w+\]`, []byte(path))
	if matched {
		handle404(w)
		return
	}

	if strings.Index(path, "/static/") != 0 {
		if path == "/" {
			files := append([]string{"./public/index.html"}, templateHome)
			t, _ := template.ParseFiles(files...)
			t.Execute(w, nil)
		} else if path != "/favicon.ico" {
			files := append([]string{fmt.Sprintf("./public%s.html", path)}, templateHome)
			t, err := template.ParseFiles(files...)
			if err == nil {
				t.Execute(w, nil)
			} else {
				fileName := strings.Split(path, "/")
				queryableValue := &fileName[len(fileName)-1]
				directory := strings.Join(fileName[:len(fileName)-1], "/")
				directoryFiles, err := ioutil.ReadDir(fmt.Sprintf("./public/%s", directory))
				if err != nil {
					handle404(w)
					return
				}
				for _, file := range directoryFiles {
					if !file.IsDir() {
						isFile, _ := regexp.Match(`\[\w+\]`, []byte(file.Name()))
						if isFile {
							jsonBytes, err := loadFile("./noSQL.json")
							if err != nil {
								handle404(w)
								return
							}

							jsonMap := map[string][]interface{}{}
							queryKey := regexp.MustCompile(`\[|\]`).Split(file.Name(), -1)[1]

							json.Unmarshal([]byte(jsonBytes), &jsonMap)

							for _, val := range jsonMap["data"] {
								for key, value := range val.(map[string]interface{}) {
									if key == queryKey && *queryableValue == value {
										fullDirectory := fmt.Sprintf("./public%s/%s", directory, file.Name())
										files := append([]string{fullDirectory}, templateHome)
										t, _ := template.ParseFiles(files...)
										fmt.Println(val)
										t.Execute(w, val)
										return
									}
								}
							}
						}
					}
				}
				handle404(w)
				return
			}
		}
	}
	return
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path != "/static/" {
		http.ServeFile(w, r, "./public"+path)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handleUri)
	http.HandleFunc("/static/", handleStatic)
	log.Fatal(http.ListenAndServe(":3010", nil))
}
