package api

import (
	"AdAlpha/db"
	"AdAlpha/logger"
	"AdAlpha/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

//Initialises endpoints
func (a *Api) InitialiseHistoryRoutes() {
	a.Router.HandleFunc("/history/investor/{id}", GetInvestorHistory).Methods("GET")
}

//Handles /history/investor/{id} call
func GetInvestorHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if id == 0 || err != nil {
		respondWithError(w, http.StatusBadRequest, "No investor id provided")
		return
	}

	dbCon := GetDbConnection()

	exists, err := db.InvestorIdExists(dbCon.Pg, id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	} else if !exists {
		respondWithError(w, http.StatusNotFound, "Investor id not found")
		return
	}

	history, err := model.GetInvestorHistory(dbCon.Pg, id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = dbCon.Pg.Close()
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR", "Error closing DB connection", err.Error(), logger.Trace(), time.Now()))
	}

	respondWithJSON(w, http.StatusOK, history)
}
