package main

import (
	"errors"
	"fmt"
	"net/http"
)

var portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello GopherWeb")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	x, y := 100.0, 10.0
	f, err := divideValues(x, y)

	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")
	} else {
		fmt.Fprintf(w, fmt.Sprintf(" %f divided by %f is %f", x, y, f))
	}
}

func divideValues(x, y float64) (float64, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	// fmt.Println("Hello Gopher!")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	_ = http.ListenAndServe(portNumber, nil)
}
