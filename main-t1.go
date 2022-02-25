package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/* also to change the keys to the json response ... change the tags against the field names */
type Customer struct {
	Name    string	`json:"full_name"`
	City    string	`json:"city"`
	Zipcode string	`json:"zip_code"`
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Alex", City: "ND", Zipcode: "1234"},
	}

	/*
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else { /* {% for json-data %} */ }
	*/

	/* encode them in json format */
	/* instruct the response writer to set the right content type response header */
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func main() {

	/* http request provides a default request multiplexer */
	/* fun-1: to register all the routes, and fun-2: to start the server */
	/* defining our routes */
	/* http.HandleFunc("route/end-point", "handlerFunc or response-writer") */
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, ".") }) /* fmt.Fprintf(i/o-writer, "data we need to send to the i/o writer") */
	http.HandleFunc("/customers", getAllCustomers)

	/* starting the server */
	log.Fatal(http.ListenAndServe("localhost:8000" /*handler*/, nil))
}

/*
func greet(w http.ResponseWriter, r *http.Request){
	fmt.Print(w, ".")
}

func main(){
	http.HandleFunc(".", greet)
}
*/

/* for xml-encoding, rep:json<->xml */
/* instead of using the default multiplexer we can also use out own-request multiplixer using the net/http lib, 
	mux := http.NewServeMux() 
	now instead of using http we can use mux to handle/register handler functions
*/
