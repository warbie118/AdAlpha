package db

import (
	"AdAlpha/logger"
	"database/sql"
	"fmt"
	"time"
)

//Checks if investor_id exists in investors table
func InvestorIdExists(db *sql.DB, id int) (bool, error) {

	var exists bool
	err := db.QueryRow("select exists(select 1 from investors where investor_id=$1)", id).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		esLog.LogError(logger.CreateLog(fmt.Sprintf("error checking if investor id exists '%d' %v", id, err),
			"Issue initialising db", err.Error(), logger.Trace(), time.Now()))

	}
	return exists, err

}
