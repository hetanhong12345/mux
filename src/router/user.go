package router

import (
	"service"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func userMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("userMiddleware--->")
		next.ServeHTTP(w, r)

	})
}

func InitUserRouter(r *mux.Router) {
	r.Use(userMiddleware)
	r.Handle("", WrapFunc(service.GetUser))
	r.Handle("/getInfo", WrapFunc(service.GetInfo))
	r.Handle("/register", WrapFunc(service.UserRegister)).Methods("POST")
}
