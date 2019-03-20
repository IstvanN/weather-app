package main

import (
	"log"
	"net/http"
	"text/template"
)

const (
	port   = ":8080"
	apiURL = "https://www.metaweather.com/api/location/"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("app is up and running on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var cityID string
	if r.Method == http.MethodPost {
		cityID = r.FormValue("cityID")
	}

	tmpl.Execute(w, cityID)
}
