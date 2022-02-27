package main

import (
	"log"
	"net/http"
	"github.com/anshugupta673/go-banking-application-dev/handlers"
)

func main() {
	//define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	/* http.HandleFunc("") */

	//starting server
	log.Fatal(http.ListenAndServe("localhost: 8000", nil))
}
