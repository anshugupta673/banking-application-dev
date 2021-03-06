package main

import(
	""
)

func start() {
	router := mux.NewRouter()

	/* wiring */
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	/* define routes */
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	/* starting server */
	log.Fatal(http.ListenAndServe("localhost:"8000", router))
}
