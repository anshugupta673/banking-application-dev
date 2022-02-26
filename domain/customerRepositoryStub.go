package domain

typr CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error){
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{Id: "1001", Name:"Alex", City: "ND", Zipcode: "1234", DateofBirthL "1234-12-12", Status: "1"}, 
		{Id: "1002", Name: "Rob", City: "ND", Zipcode: "110011", DateofBirth: "2012-12-12", Status: "1"}
	}
	return CustomerRepositoryStub{ customers: customers }
}
