package db

import (
	"log"
)

func init() {
	connectDb()
}

func ExecuteSql(sqlReq string) ([]string, error) {
	rows, err := db.Query(sqlReq)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	result := make([]string, 0)
	for rows.Next() {
		var id, text string
		if err := rows.Scan(&id, &text); err != nil {
			log.Fatal(err)
		}
		result = append(result, id+text)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result, err
}
