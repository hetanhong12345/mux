package user

import (
	"sql"
	"models"
	"errors"
)

var db = sql.DB()

type User = models.User

func Register(mobile, password string) (*User, error) {
	user := User{Mobile: mobile, Password: password}
	if err := db.Where(&User{Mobile: mobile}).First(&user).Error; err == nil {
		return &User{}, errors.New("mobile has registed")
	}
	if err := db.Create(&user).Error; err != nil {
		return &User{}, err
	}
	return &user, nil

}

func Login(mobile, password string) (*User, error) {
	user := User{Mobile: mobile, Password: password}
	if err := db.Where(&User{Mobile: mobile}).First(&user).Error; err != nil {
		return &User{}, err
	}
	if user.Password != password {
		return &User{}, errors.New("password is incorrect")
	}
	return &user, nil
}
