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

type CustomerHandlers struct{
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := 
}
