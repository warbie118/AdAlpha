package db

import (
	"AdAlpha/logger"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

type Db struct {
	Pg *sql.DB
}

var esLog = logger.GetInstance()

//Initialises DB connection
func (db *Db) Initialise() {
	connectionString :=
		fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", "localhost", os.Getenv("DB_PORT"),
			os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), "disable")
	var err error
	db.Pg, err = sql.Open("postgres", connectionString)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR", "Issue initialising db", err.Error(), logger.Trace(), time.Now()))
		log.Fatal(err)
	}
}
