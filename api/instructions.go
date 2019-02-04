package api

import (
	"fmt"
	"net/http"
)

//Initialises endpoints
func (a *Api) InitialiseInstructionRoutes() {
	a.Router.HandleFunc("/instruction/buy", NewBuy).Methods("POST")
	a.Router.HandleFunc("/instruction/invest", NewInvest).Methods("POST")
	a.Router.HandleFunc("/instruction/sell", NewSell).Methods("POST")
	a.Router.HandleFunc("/instruction/raise", NewRaise).Methods("POST")
	a.Router.HandleFunc("/history/investor/{id}", GetInvestorHistory).Methods("GET")
}

//Handles /instruction/buy call
func NewBuy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("not implemented")
}

//Handles /instruction/invest call
func NewInvest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("not implemented")
}

//Handles /instruction/raise call
func NewRaise(w http.ResponseWriter, r *http.Request) {
	fmt.Println("not implemented")
}

//Handles /instruction/sell call
func NewSell(w http.ResponseWriter, r *http.Request) {
	fmt.Println("not implemented")
}
