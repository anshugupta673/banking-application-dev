package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

/* also to change the keys to the json response ... change the tags against the field names */
type Customer struct {
	Name    string `json: "full_name" xml:name`
	City    string `json: "city" xml: "city"`
	Zipcode string `json: "zip_code" xml: "zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, ".")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Alex", City: "ND", Zipcode: "123"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
