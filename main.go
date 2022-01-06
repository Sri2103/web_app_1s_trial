package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const port string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("this is a Home Page"))

}

func About(w http.ResponseWriter, r *http.Request) {

	sum := AddItem(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("the  value of 2+2 is %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, fmt.Sprintf("this is a Home Page"))
	f, err := divideNumbers(100.00, 0.00)

	if err != nil {
		fmt.Fprintf(w, "cannot be divide by 0. ")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f gives the value of %f", 100.00, 0.00, f))

}

func AddItem(x, y int) int {
	return x + y
}

func divideNumbers(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot be divided by zero")

		return 0, err

	}

	result := x / y

	return result, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	log.Printf("listening on %s", port)

	_ = http.ListenAndServe(port, nil)
}
