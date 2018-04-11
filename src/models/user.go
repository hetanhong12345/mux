package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Uuid string `gorm:"not null;unique" json:"uuid"` // 设置字段为非空并唯一
	Mobile string `gorm:"not null;unique" json:"mobile"` // 设置字段为非空并唯一
	Password string `gorm:"not null" json:"password"`

}
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	u2, err := uuid.NewV4()
	if err != nil {
		panic("uuid error")
	}
	scope.SetColumn("Uuid", u2.String())
	return nil
}