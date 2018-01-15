package common

import "fmt"

func NewLoginer()  Loginer{
	return  defaultLogin(100)
}

type  Loginer interface {
	Login()
}

type defaultLogin int

func (d defaultLogin) Login()  {
	fmt.Printf("login in\n")
	fmt.Printf("%d\n",d)
}