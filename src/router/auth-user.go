package router

import (
	"controllers/user"
	"middlewares"
	"github.com/gorilla/mux"
)

func InitAUthUserRouter(r *mux.Router) {

	r.Use(middlewares.LoginRequried)
	r.HandleFunc("/changeName", user.ChangeName).Methods("POST")

}