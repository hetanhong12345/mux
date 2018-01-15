package main

import "fmt"

type int1 int
type int2 = int

func main() {
	var i int = 0
	var i1 int1 = 4
	var i2 int2 = i
	fmt.Println(i1, i2)
	c := add(2)(4)
	fmt.Println(c)
}

func add(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}

}
