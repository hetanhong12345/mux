// hello
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var c = 2

	for c < 20 {
		c++
		fmt.Printf("faibo %d is %d \n", c, faibo(c))
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(time.Second * 1)
		for i := 1; i < 3; i++ {
			fmt.Println("A:", i)
		}
		defer wg.Done()
	}()
	go func() {

		for i := 1; i < 3; i++ {
			fmt.Println("B:", i)
		}
		time.Sleep(time.Second * 2)
		fmt.Println("defer done")
		defer wg.Done()

	}()
	wg.Wait()
	fmt.Println("it's ok")
	one := make(chan int)
	two := make(chan int)

	go func() {
		fmt.Println("---->two start")
		v := <-one
		two <- v
		fmt.Println("---->two end")

	}()
	go func() {
		fmt.Println("---->one start")
		time.Sleep(time.Second * 2)
		one <- 100
		fmt.Println("---->one end")

	}()
	fmt.Println("<-----end")
	twoValue := <-two
	fmt.Println("two", twoValue)
	fmt.Println("<-----end---->")
}
func faibo(num int) int {
	if num <= 2 {
		return 1
	}
	return faibo(num-1) + faibo(num-2)
}
