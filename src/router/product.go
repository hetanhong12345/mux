package router

import (
	"service"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func productMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("productMiddleware--->")
		next.ServeHTTP(w, r)

	})
}

func InitProductRouter(r *mux.Router) {

	r.Use(productMiddleware)

	r.Handle("", WrapFunc(service.GetProduct)).Methods("GET")
	r.Handle("/info", WrapFunc(service.GetProductInfo)).Methods("GET")
}
