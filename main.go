package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const port string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")

}

func renderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error in Parsling Template :", err)
		return
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("listening on %s", port)

	_ = http.ListenAndServe(port, nil)
}
