package api

import (
	"fmt"
	"net/http"
)

//Initialises endpoints
func (a *Api) InitialiseHistoryRoutes() {
	a.Router.HandleFunc("/history/investor/{id}", GetInvestorHistory).Methods("GET")
}

//Handles /history/investor/{id} call
func GetInvestorHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("not implemented")
}
