package service

import (
	"net/http"
	"fmt"
)

func GetProduct(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(rw, "request products")
}
func GetProductInfo(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(rw, "request productsInfo")
}