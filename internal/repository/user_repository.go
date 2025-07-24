package repository

import (
	"fmt"
	"log"
	"userTest/config"
)

func ListUsers() {
	rows, err := config.DB.Query("SELECT id, username FROM users")
	if err != nil {
		log.Println("ERROR: COULDN'T READ USERS:", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var username string

		err := rows.Scan(&id, &username)
		if err != nil {
			log.Println("ERROR: COULDN'T LIST ANY USER:", err)
			continue
		}
		fmt.Printf("\n|ID: %d | USERNAME: %s |", id, username)
	}

	if err = rows.Err(); err != nil {
		log.Println("ERROR: COULDN'T ITERATE THROUGHT USERS", err)
	}
}
