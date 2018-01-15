package service

import (
	"net/http"
	"fmt"
)

func GetUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "request getUser")
}
func GetInfo(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "request getUserInfo")
}