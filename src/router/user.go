package router

import (
	"service"
)

func init() {
	User := Routes.PathPrefix("/user").Subrouter()

	User.Handle("/", WrapFunc(service.GetUser))
	User.Handle("/getInfo", WrapFunc(service.GetInfo))
	User.Handle("/register", WrapFunc(service.UserRegister)).Methods("POST")

}
