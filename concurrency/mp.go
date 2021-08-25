package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	tiempo1 := time.Second * 4
	tiempo2 := time.Second * 2

	go DoSomething(tiempo1, c1, 1)
	go DoSomething(tiempo2, c2, 2)

	// fmt.Println(<-c1)
	// fmt.Println(<-c2)

	for i := 0; i < 2; i++ {
		select {
		case message1 := <-c1:
			fmt.Println(message1)
		case message2 := <-c2:
			fmt.Println(message2)
		}
	}

}
func DoSomething(t time.Duration, c chan<- int, param int) {
	time.Sleep(t)

	c <- param
}
