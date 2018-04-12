package router

import (
	"github.com/gorilla/mux"
	"middlewares"
)

var RootRouter *mux.Router = mux.NewRouter()
var UserRouter *mux.Router

func init() {
	RootRouter.Use(middlewares.Logger)
	RootRouter.Use(middlewares.ResponseHeader)
	UserRouter = RootRouter.PathPrefix("/user").Subrouter()
	InitUserRouter(UserRouter)
}

