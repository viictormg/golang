package main

import (
	"fmt"
	"time"
)

func main1() {
	// x := 5

	// y := func() int {
	// 	return x * 2
	// }()
	// fmt.Println(y)

	c := make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(5 * time.Second)
		c <- 1
		fmt.Println("End")
	}()

	fmt.Println(<-c)
}
