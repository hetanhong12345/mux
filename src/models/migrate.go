package models

import (
	"sql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {

	db = sql.DB()

	db.AutoMigrate(&User{})

}
