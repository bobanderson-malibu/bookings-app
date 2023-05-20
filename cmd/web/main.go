package main

import (
	"net/http"

	"github.com/bobanderson-malibu/bookings-app/pkg/handlers"
)

var portNumber = ":8080"

func main() {
	// fmt.Println("Hello Gopher!")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)
}
