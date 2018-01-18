package router

import (
	"github.com/gorilla/mux"

	"net/http"
	"fmt"
)

var Routes *mux.Router = mux.NewRouter()


/*
 params  http.HandlerFunc
 return  http.Handler
 description http中间件

*/

func WrapFunc(handlerFunc http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("<----middleware---WrapFunc-->")
		//fmt.Println(r.Header.Get("User-Agent"))
		handlerFunc(rw, r)

	})
}