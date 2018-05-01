package router

import (
	"controllers/user"
	"github.com/gorilla/mux"
)

func InitUserRouter(r *mux.Router) {

	r.HandleFunc("/register", user.Register).Methods("POST")
	r.HandleFunc("/login", user.Login).Methods("POST")

}
