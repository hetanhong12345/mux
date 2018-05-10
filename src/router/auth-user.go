package router

import (
	"controllers/user"
	"github.com/gorilla/mux"
	"middlewares"
)

func InitAUthUserRouter(r *mux.Router) {

	r.Use(middlewares.LoginRequired)
	r.HandleFunc("/changeName", user.ChangeName).Methods("POST")

}
