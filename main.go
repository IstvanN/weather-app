package main

import (
	"encoding/json"
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
	if r.Method == http.MethodPost {
		cityID := r.FormValue("cityID")
		resp, err := http.Get(apiURL + cityID)
		if err != nil {
			log.Println(err)
		}

		var rd rawData
		json.NewDecoder(resp.Body).Decode(&rd)
		wd = getHighestPredict(rd)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl.Execute(w, wd)
	wd = WeatherData{}
}
