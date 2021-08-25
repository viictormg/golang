package main

import (
	"fmt"
	"sync"
	"time"
)

func main4() {
	c := make(chan int, 4)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- i
		wg.Add(1)
		go doSomething2(i, &wg, c)
	}

	wg.Wait()
}

func doSomething2(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Started %d \n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Finished %d \n", i)
	<-c

}
