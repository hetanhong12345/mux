package router

import (
	"github.com/gorilla/mux"
	"middlewares"
)

var Root *mux.Router

func init() {
	Root = mux.NewRouter()
	Root.Use(middlewares.Logger)
	Root.Use(middlewares.ResponseHeader)

	User := Root.PathPrefix("/user").Subrouter()
	AuthUser := Root.PathPrefix("/auth-user").Subrouter()
	InitUserRouter(User)
	InitAUthUserRouter(AuthUser)
}
