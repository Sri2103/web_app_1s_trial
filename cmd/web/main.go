package main

import (
	"log"
	"myapp/pkg/headers"
	"net/http"
)

const port string = ":8080"

func main() {

	http.HandleFunc("/", headers.Home)
	http.HandleFunc("/about", headers.About)

	log.Printf("listening on %s", port)

	_ = http.ListenAndServe(port, nil)
}
