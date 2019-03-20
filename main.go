package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("app is up and running on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello weather!")
}
