package main

import "fmt"

func main2() {
	c := make(chan int, 2)

	c <- 1

	fmt.Println(<-c)
}
