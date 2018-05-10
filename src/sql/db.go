package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sync"
)

var db *gorm.DB
var err error
var loadDBOnce sync.Once

// export db
func DB() *gorm.DB {

	loadDBOnce.Do(initBD)
	log.Println("load once sql.DB()")

	return db

}

func initBD() {
	db, err = gorm.Open("mysql", "root:hxx123456@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		panic("mysql connect err")

	}
	log.Println("init sql.DB()")
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
}
