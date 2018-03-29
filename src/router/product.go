package router

import (
	"service"
	"net/http"
	"fmt"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middlewareA--->")
		next.ServeHTTP(w, r)

	})
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		fmt.Printf("url is %s ----->\n", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
func init() {

	Product := Routes.PathPrefix("/product").Subrouter()
	Product.Use(loggingMiddleware)
	Product.Use(middleware)

	Product.Handle("/", WrapFunc(service.GetProduct)).Methods("GET")
	Product.Handle("/info", WrapFunc(service.GetProductInfo)).Methods("GET")

}
