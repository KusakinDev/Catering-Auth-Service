package main

import (
	"log"

	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
)

func main() {
	var db database.DataBase
	db.InitDB()

	sqlStatements := []string{
		`INSERT INTO roles (role, role_string) VALUES
        ('adm', 'Администратор'),
        ('mng', 'Менеджер'),
        ('wtr', 'Официант'),
        ('ktn', 'Кухня'),
        ('bar', 'Бар');`,

		`INSERT INTO user_accounts (id, email, password, role_id) VALUES
        (1, 'user1@example.com', 'password123', 1),
        (2, 'mod1@example.com', 'password123', 2),
        (3, 'admin1@example.com', 'password123', 3),
        (4, 'user2@example.com', 'password456', 4),
        (5, 'mod2@example.com', 'password456', 5);`,
	}

	for _, stmt := range sqlStatements {
		if err := db.Connection.Exec(stmt).Error; err != nil {
			log.Println("Error executing seed: ", stmt, err)
		}
	}

	log.Println("Success seeding")

	db.CloseDB()
}
