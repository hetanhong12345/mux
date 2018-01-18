package service

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func GetUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "request getUser")
}
func GetInfo(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "request getUserInfo")
}
func UserRegister(rw http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fromData := fmt.Sprintf("%#v", r.Form)
	fmt.Println(fromData)
	fmt.Println(r.Form["mobile"][0])
	account := Account{
		Email:    "rsj217@gmail.com",
		Password: "123456",
		Money:    100.5,
	}
	outJson, _ := json.Marshal(account)
	rw.Write(outJson);
}
