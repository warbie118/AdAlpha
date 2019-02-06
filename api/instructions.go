package api

import (
	"AdAlpha/model"
	"encoding/json"
	"net/http"
)

//Initialises endpoints
func (a *Api) InitialiseInstructionRoutes() {
	a.Router.HandleFunc("/instruction/buy", NewBuy).Methods("POST")
	a.Router.HandleFunc("/instruction/invest", NewInvest).Methods("POST")
	a.Router.HandleFunc("/instruction/sell", NewSell).Methods("POST")
	a.Router.HandleFunc("/instruction/raise", NewRaise).Methods("POST")
}

//Handles /instruction/buy call
func NewBuy(w http.ResponseWriter, r *http.Request) {
	var buyReq model.Buy

	if r.Body == nil {
		respondWithError(w, http.StatusBadRequest, "Please send a request body")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&buyReq); err != nil || !buyReq.IsValid() {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	//implement calculate func /model/buy
	buyReq.CalculateBuy()

	respondWithJSON(w, http.StatusOK, buyReq)

}

//Handles /instruction/invest call
func NewInvest(w http.ResponseWriter, r *http.Request) {
	var investReq model.Invest

	if r.Body == nil {
		respondWithError(w, http.StatusBadRequest, "Please send a request body")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&investReq); err != nil || !investReq.IsValid() {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	investReq.CalculateInvest()

	respondWithJSON(w, http.StatusOK, investReq)

}

//Handles /instruction/raise call
func NewRaise(w http.ResponseWriter, r *http.Request) {
	var raiseReq model.Raise
	if r.Body == nil {
		respondWithError(w, http.StatusBadRequest, "Please send a request body")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&raiseReq); err != nil || !raiseReq.IsValid() {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	raiseReq.CalculateRaise()

	respondWithJSON(w, http.StatusOK, raiseReq)

}

//Handles /instruction/sell call
func NewSell(w http.ResponseWriter, r *http.Request) {
	var sellReq model.Sell
	if r.Body == nil {
		respondWithError(w, http.StatusBadRequest, "Please send a request body")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&sellReq); err != nil || !sellReq.IsValid() {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	sellReq.CalculateSell()

	respondWithJSON(w, http.StatusOK, sellReq)

}
