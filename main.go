package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func uriHandler(res http.ResponseWriter, req *http.Request) {
	requestUri := req.URL.Path
	if requestUri == "/" {
		requestUri = "./public/index.html"
	} else {
		requestUri = fmt.Sprintf("./public/%s.html", requestUri)
	}

	if req.Method == "GET" {
		requestedFile, err := loadFile(requestUri)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			requestedFile, err := loadFile("./public/404.html")
			if err != nil {
				fmt.Fprintf(res, "Some dipshit didn't include a 404 page. Either way, your page wasn't found.")
			}
			fmt.Fprintf(res, requestedFile)
		} else {
			res.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(res, requestedFile)
		}
	} else {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "%s is not an accepted Method", req.Method)
	}
}

func main() {
	http.HandleFunc("/", uriHandler)
	http.ListenAndServe(":3010", nil)
}
