package api

import "github.com/gorilla/mux"

type Api struct {
	Router *mux.Router
}

//Calls to instructions.go and history.go to initialise api routes
func (a *Api) Initialise() {
	a.Router = mux.NewRouter()
	a.InitialiseInstructionRoutes()
	a.InitialiseHistoryRoutes()
}
