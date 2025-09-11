package main

import (
	"fmt"
	"sync"
)
func GetNumbers(numbers []int,out chan<- int){
	for _,i := range numbers{
		out<-i
	}
	close(out)
}
func Square(in <-chan int, out chan<- int){
	for i:= range in {
		out<-i*i
	}
close(out)
}
func Printer(in<-chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for i:=range in{
		fmt.Printf("%v\n",i)
	}
}
func main() {
	jobs := make(chan int)
	results :=make(chan int)
	var wg sync.WaitGroup
	wg.Add(1) 
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	go GetNumbers(nums, jobs)
	go Square(jobs,results)
	go Printer(results,&wg)
	wg.Wait()
}
