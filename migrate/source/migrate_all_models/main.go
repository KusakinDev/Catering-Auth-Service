package main

import (
	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
	accountmodel "github.com/KusakinDev/Catering-Auth-Service/internal/models/account_model"
	resetpasswordmodel "github.com/KusakinDev/Catering-Auth-Service/internal/models/reset_password_model"
)

func main() {

	var db database.DataBase
	db.InitDB()

	var account accountmodel.UserAccount
	account.MigrateToDB(db)

	var resetCode resetpasswordmodel.ResetCode
	resetCode.MigrateToDB(db)

	db.CloseDB()
}
