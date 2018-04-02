package router

import (
	"github.com/gorilla/mux"

	"net/http"
	"fmt"
)

var RootRouter *mux.Router = mux.NewRouter()

func init() {
	RootRouter.Use(loggingMiddleware)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("<----log middleware request url is %s ----->\n", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

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
