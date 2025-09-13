//Бедрин совместно с Жироховым
package main

import (
	"fmt"
	"sync"
)

var jobs = make(chan int)
var result = make(chan int)
var wg sync.WaitGroup

func Worker(jobs <-chan int, result chan<- int) {
	defer wg.Done()
	for arg := range jobs {
		result <-arg
	}
}

func main() {

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Worker(jobs, result)
	}


	go func() {
		for i:=0; i<100;i++{
			jobs <- i
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for i := range result {
		fmt.Printf("%v\n", i)
	}
}
