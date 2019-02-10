package api

import (
	"AdAlpha/db"
	"AdAlpha/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var esLog = logger.GetInstance()

type Api struct {
	Router *mux.Router
}

//Calls to instructions.go and history.go to initialise api routes
func (a *Api) Initialise() {
	a.Router = mux.NewRouter()
	a.InitialiseInstructionRoutes()
	a.InitialiseHistoryRoutes()
	a.InitialisePortfolioRoutes()
}

//create a connection to DB
func GetDbConnection() db.Db {
	d := db.Db{}
	d.Initialise()
	return d
}

//returns json to the api request
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}

//calls respondWithJson to return error to api request
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
