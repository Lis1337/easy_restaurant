package db

import (
	"database/sql"
	"log"
)

func ExecuteSql(sql string, args ...any) (*sql.Rows, error) {
	connectDb()

	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		log.Fatal(err)
	}

	defer closeCon()

	return rows, err
}
