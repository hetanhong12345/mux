package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/satori/go.uuid"
)

type Kingold struct {
	gorm.Model
	Name string
	Uuid string `gorm:"not null;unique"` // 设置字段为非空并唯一
	Age  int8
}
type Student struct {
	gorm.Model
	Skill     string
	Expirence int8
	ClassNum  int8
}
type Teacher struct {
	gorm.Model
	Name string `gorm:"not null"`
	Age  int8   `gorm:"not null"`
}

func (kingold *Kingold) BeforeCreate(scope *gorm.Scope) error {
	u2, err := uuid.NewV4()
	if err != nil {
		panic("uuid error")
	}
	scope.SetColumn("Uuid", u2.String())
	return nil
}
func (teacher *Teacher) BeforeCreate(scope *gorm.Scope) error {
	if teacher.Age == 0 {
		panic("Age can not be 0")
	}
	return nil
}

func main() {
	db, err := gorm.Open("mysql", "root:hxx123456@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("mysql connect err")

	}
	defer db.Close()
	student := Student{Skill: "c++", Expirence: 6}
	db.Create(&student)
	db.AutoMigrate(&Kingold{})
	teacher := Teacher{Name: "fujiajun", Age: 36}
	db.Create(&teacher)
	kingold := Kingold{Name: "qiqi", Age: 18}
	db.Debug().Create(&kingold)
	kingolds := []Kingold{}
	db.Debug().Find(&kingolds)

}
