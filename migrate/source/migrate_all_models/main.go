package main

import (
	"log"

	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
	accountmodel "github.com/KusakinDev/Catering-Auth-Service/internal/models/account_model"
	resetpasswordmodel "github.com/KusakinDev/Catering-Auth-Service/internal/models/reset_password_model"
	rolemodel "github.com/KusakinDev/Catering-Auth-Service/internal/models/role_model"
)

func main() {

	var db database.DataBase
	db.InitDB()

	var account accountmodel.UserAccount
	account.MigrateToDB(db)

	var role rolemodel.Role
	role.MigrateToDB(db)

	var resetCode resetpasswordmodel.ResetCode
	resetCode.MigrateToDB(db)

	sqlStatements := []string{
		`INSERT INTO roles (role, role_string) VALUES
        ('adm', 'Администратор'),
        ('mng', 'Менеджер'),
        ('wtr', 'Официант'),
        ('ktn', 'Кухня'),
        ('bar', 'Бар');`,
	}

	for _, stmt := range sqlStatements {
		if err := db.Connection.Exec(stmt).Error; err != nil {
			log.Println("Error executing seed: ", stmt, err)
		}
	}

	log.Println("Success seeding")

	db.CloseDB()
}
