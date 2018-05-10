package models

import (
	"github.com/jinzhu/gorm"
	"sql"
)

var db *gorm.DB

func init() {

	db = sql.DB()

	db.AutoMigrate(&User{})

}
