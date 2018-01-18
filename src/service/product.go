package service

import (
	"net/http"
	"fmt"
	"encoding/json"
	"time"
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
	rw.Write(outJson);
}
func GetProductInfo(rw http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(3) * time.Second)

	fmt.Fprintf(rw, "request productsInfo")
}
