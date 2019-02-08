package db

import (
	"database/sql"
	"log"
)

//Checks if investor_id exists in investors table
func InvestorIdExists(db *sql.DB, id int) (bool, error) {

	var exists bool
	err := db.QueryRow("select exists(select 1 from investors where investor_id=$1)", id).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("error checking if investor id exists '%d' %v", id, err)
	}
	return exists, err

}
