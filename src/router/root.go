package router

import (
	"github.com/gorilla/mux"
	"middlewares"
)

var Root *mux.Router = mux.NewRouter()
var User *mux.Router
var AuthUser *mux.Router

func init() {
	Root.Use(middlewares.Logger)
	Root.Use(middlewares.ResponseHeader)
	User = Root.PathPrefix("/user").Subrouter()
	AuthUser = Root.PathPrefix("/auth-user").Subrouter()
	InitUserRouter(User)
	InitAUthUserRouter(AuthUser)
}
