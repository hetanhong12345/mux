package router

import (
	"service"
)

func init() {
	Product := Routes.PathPrefix("/product").Subrouter()

	Product.Handle("/", WrapFunc(service.GetProduct)).Methods("GET")
	Product.Handle("/info", WrapFunc(service.GetProductInfo)).Methods("GET")

}
