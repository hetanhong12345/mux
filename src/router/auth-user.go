package router

import (
	"controllers/user"
	"middlewares"
	"github.com/gorilla/mux"
)

func InitAUthUserRouter(r *mux.Router) {

	r.Use(middlewares.LoginRequired)
	r.HandleFunc("/changeName", user.ChangeName).Methods("POST")

}