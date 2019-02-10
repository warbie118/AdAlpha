package main

import (
	"AdAlpha/api"
	"AdAlpha/db"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {

	a := api.Api{}
	d := db.Db{}
	a.Initialise()
	d.Initialise()
	handler := cors.Default().Handler(a.Router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
