package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"models"
	"fmt"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("mysql", "root:hxx123456@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("mysql connect err")

	}
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	migrate()
}

// export db
func DB() *gorm.DB {
	return db

}

// migrate tables
func migrate() {
	DB().AutoMigrate(&models.User{})

}
