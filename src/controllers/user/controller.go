package user

import (
	"net/http"
	"fmt"
	"middlewares"
	"services/user"
)

type Result struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	result := &Result{}

	defer middlewares.Response(w, result)
	mobile := r.PostForm["mobile"][0]
	password := r.PostForm["password"][0]
	if (mobile == "") {
		result.Code = "400"
		result.Msg = "mobile is empty"
		return

	}
	if (password == "") {
		result.Code = "400"
		result.Msg = "password is empty"
		return

	}
	userInfo, err := user.Register(mobile, password)
	if err != nil {
		result.Code = "500"
		result.Msg = fmt.Sprintf("%v", err)
		return
	}
	result.Code = "200"
	result.Msg = "ok"
	result.Data = userInfo

}
func Login(w http.ResponseWriter, r *http.Request) {
	result := &Result{}

	defer middlewares.Response(w, result)
	mobile := r.PostForm["mobile"][0]
	password := r.PostForm["password"][0]
	if (mobile == "") {
		result.Code = "400"
		result.Msg = "mobile is empty"
		return

	}
	if (password == "") {
		result.Code = "400"
		result.Msg = "password is empty"
		return

	}
	userInfo, err := user.Login(mobile, password)
	if err != nil {
		result.Code = "500"
		result.Msg = fmt.Sprintf("%v", err)
		return
	}
	result.Code = "200"
	result.Msg = "ok"
	result.Data = userInfo
	cookie := http.Cookie{Name: "mobile", Value: mobile, Path: "/", MaxAge: 86400}
	http.SetCookie(w, &cookie)

}

func ChangeName(w http.ResponseWriter, r *http.Request) {
	result := &Result{}
	defer middlewares.Response(w, result)
	cookie, _ := r.Cookie("mobile")
	name := r.PostForm["name"][0]
	mobile := cookie.Value
	if name == "" {
		result.Code = "400"
		result.Msg = "name is empty"
		return
	}
	userInfo, err := user.ChangeName(mobile, name)
	if err != nil {
		result.Code = "500"
		result.Msg = fmt.Sprintf("%v", err)
		return
	}
	result.Code = "200"
	result.Msg = "ok"
	result.Data = userInfo
}
