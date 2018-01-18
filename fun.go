package main

import (
	"encoding/json"
	"log"
	"fmt"
	"time"
)

type User struct {
	Name  string
	Age   int
	Roles []string
	Skill map[string]float64
}
type Account struct {
	Name     string `json:"name"`
	Password string `json:"-"`
}

func main() {
	account := Account{"hth", "hxx123456"}
	accountJson, err := json.Marshal(account)
	if (err != nil) {
		log.Fatalln(err)
	}
	fmt.Println(string(accountJson))
	skill := make(map[string]float64)

	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0
	myJson := make(map[string]interface{})
	myJson["aa"] = "bb"
	myJson["12"] = "34"
	myJson["_ad"] = true
	fmt.Println(myJson)

	user := User{
		Name:  "rsj217",
		Age:   27,
		Roles: []string{"Owner", "Master"},
		Skill: skill,
	}

	rs, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	outJson, err := json.Marshal(myJson)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(outJson))

	fmt.Println(string(rs))
	fmt.Println(time.Now())
}
