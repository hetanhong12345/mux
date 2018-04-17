package models

import "sql"

var db = sql.DB()

func init() {

	db.AutoMigrate(&User{})

}
