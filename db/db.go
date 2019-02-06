package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Db struct {
	Pg *sql.DB
}

//Initialises DB connection
func (db *Db) Initialise() {
	connectionString :=
		fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", "localhost", os.Getenv("DB_PORT"),
			os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), "disable")
	var err error
	db.Pg, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}
