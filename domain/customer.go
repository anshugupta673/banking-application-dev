package domain

import (
	"fmt"
)

type Customer struct {
	Id string
	Name city
	Zipcode string
	DateofBirth string
	Status string
}

/* repository interface */
type CustomerRepositiry interface { /* finding all the customers from the server-side */
	FindAll() ([]Customer, error)
}
/*
/* stub-adapter */
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error){
	return s.customers, nil
}

func New
 
*/
