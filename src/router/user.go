package router

import (
	"fmt"
	"service"
)

func init() {
	User := Routes.PathPrefix("/user").Subrouter()

	User.Handle("/", WrapFunc(service.GetUser))
	User.Handle("/getInfo", WrapFunc(service.GetInfo))
	fmt.Println("user")

}

