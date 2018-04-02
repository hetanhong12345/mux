package router

import (
	"service"
	"net/http"
	"fmt"
)

func userMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("userMiddleware--->")
		next.ServeHTTP(w, r)

	})
}
func init() {
	SubRouter := RootRouter.PathPrefix("/user").Subrouter()
	SubRouter.Use(userMiddleware)
	SubRouter.Handle("", WrapFunc(service.GetUser))
	SubRouter.Handle("/getInfo", WrapFunc(service.GetInfo))
	SubRouter.Handle("/register", WrapFunc(service.UserRegister)).Methods("POST")

}
