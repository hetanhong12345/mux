package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"sync"
)

var db *gorm.DB
var err error
var loadDBOnce sync.Once

// export db
func DB() *gorm.DB {

	loadDBOnce.Do(initBD)

	return db

}

func initBD() {
	db, err = gorm.Open("mysql", "root:hxx123456@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("mysql connect err")

	}
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
}


