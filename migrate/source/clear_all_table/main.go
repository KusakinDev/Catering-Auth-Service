package main

import (
	"log"

	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
)

func main() {
	var db database.DataBase
	db.InitDB()

	query := []string{
		`DELETE FROM user_accounts;`,
		`DELETE FROM roles;`,
		`DELETE FROM reset_codes;`,
	}

	for _, stmt := range query {
		if err := db.Connection.Exec(stmt).Error; err != nil {
			log.Println("Error executing clear: ", stmt, err)
		}
	}

	log.Println("All table is clear")

	db.CloseDB()
}
