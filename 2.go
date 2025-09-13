package main

import (
	"fmt"
	"sync"
)

func GetNumbers(numbers []int, jobs chan<- int) {

	for _, i := range numbers {
		jobs <- i
	}
	defer close(jobs)
}
func Square(jobs <-chan int, results chan<- int) {
	for i := range jobs {
		results <- i * i
	}
}
func Printer(results <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range results {
		fmt.Printf("%v\n", i)
	}
}
func main() {
	jobs := make(chan int)
	results := make(chan int)
	var wgWorkers sync.WaitGroup
	var wgPrinter sync.WaitGroup
	numWorkers := 3
	wgWorkers.Add(numWorkers)
	wgPrinter.Add(1)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go GetNumbers(nums, jobs)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wgWorkers.Done()
			Square(jobs, results)
		}()
	}
	go func() {
		wgWorkers.Wait()
		close(results)
	}()
	go Printer(results, &wgPrinter)
	wgPrinter.Wait()
}
