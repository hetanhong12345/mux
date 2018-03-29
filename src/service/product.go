package service

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Account struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Money    float64 `json:"money"`
}

func GetProduct(rw http.ResponseWriter, r *http.Request) {
	account := Account{
		Email:    "rsj217@gmail.com",
		Password: "123456",
		Money:    100.5,
	}
	outJson, _ := json.Marshal(account)
	array := [3]int{1, 2, 3}
	index :=4
	fmt.Println(array[index])
	rw.Write(outJson);
}
func GetProductInfo(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(rw, "request productsInfo")
}
