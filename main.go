package main

import (
	"AdAlpha/api"
	"log"
	"net/http"
)

// main function - initialises API
func main() {

	a := api.Api{}
	a.Initialise()
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
