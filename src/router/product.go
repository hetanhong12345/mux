package router

import (
	"service"
	"net/http"
	"fmt"
)

func productMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("productMiddleware--->")
		next.ServeHTTP(w, r)

	})
}
func init() {

	SubRouter := RootRouter.PathPrefix("/product").Subrouter()

	SubRouter.Use(productMiddleware)

	SubRouter.Handle("", WrapFunc(service.GetProduct)).Methods("GET")
	SubRouter.Handle("/info", WrapFunc(service.GetProductInfo)).Methods("GET")

}
