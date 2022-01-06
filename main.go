package main

import (
	"log"
	"net/http"
)

const port string = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("listening on %s", port)

	_ = http.ListenAndServe(port, nil)
}
