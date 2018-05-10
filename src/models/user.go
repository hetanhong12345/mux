package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"util"
)

type User struct {
	Base
	Name     string `json:"name"`
	Uuid     string `gorm:"not null;unique" json:"uuid"`   // 设置字段为非空并唯一
	Mobile   string `gorm:"not null;unique" json:"mobile"` // 设置字段为非空并唯一
	Password string `gorm:"not null" json:"-"`
}

// 创建钱添加uuid
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	u2, err := uuid.NewV4()
	if err != nil {
		panic("uuid error")
	}

	password := util.Encrypt(user.Password)
	user.Password = password
	scope.SetColumn("Uuid", u2.String())
	return nil
}

// 添加一条记录
func (user *User) CreateOne() error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil

}

// 根据mobile查找user
func (user *User) FindByMobile(mobile string) error {
	if err := db.Where(&User{Mobile: mobile}).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// 更新部分字段
func (user *User) Update(attrs map[string]interface{}) error {
	if err := db.Model(user).Update(attrs).Error; err != nil {
		return err
	}
	return nil
}
