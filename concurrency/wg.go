package main

import (
	"fmt"
	"sync"
	"time"
)

func main3() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}
	wg.Wait()
}

func doSomething(iteration int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Started %d \n", iteration)
	time.Sleep(2 * time.Second)
	fmt.Printf("Finish %d \n", iteration)

}
