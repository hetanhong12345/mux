package user

import (
	"errors"
	"models"
	"util"
)

type User = models.User

func Register(mobile, password string) (*User, error) {
	user := User{Mobile: mobile, Password: password}
	if err := user.FindByMobile(mobile); err == nil {
		return &User{}, errors.New("mobile has registed")
	}
	if err := user.CreateOne(); err != nil {
		return &User{}, err
	}
	return &user, nil

}

func Login(mobile, password string) (*User, error) {
	user := User{}
	if err := user.FindByMobile(mobile); err != nil {
		return &User{}, err
	}
	if user.Password != util.Encrypt(password) {
		return &User{}, errors.New("password is incorrect")
	}
	return &user, nil
}

func ChangeName(mobile, name string) (*User, error) {
	user := User{}
	if err := user.FindByMobile(mobile); err != nil {
		return &User{}, errors.New("mobile is not registed")
	}
	attrs := make(map[string]interface{})
	attrs["name"] = name
	if err := user.Update(attrs); err != nil {
		return &User{}, err
	}
	return &user, nil

}
