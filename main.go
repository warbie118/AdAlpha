package main

import (
	"AdAlpha/api"
	"AdAlpha/db"
	"log"
	"net/http"
)

func main() {

	a := api.Api{}
	d := db.Db{}
	a.Initialise()
	d.Initialise()
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
