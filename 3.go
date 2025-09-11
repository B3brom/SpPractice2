package main

import (
	"fmt"
	"sync"
	"time"
)

var jobs = make(chan int,15)
var result = make(chan int,15)
var wg sync.WaitGroup

func Semaftor(){
	tick:= time.Tick(200*time.Millisecond)
	wg.Add(1)
	go func(){
		defer wg.Done()
		for i := range jobs{
			<-tick
			result <-i
			fmt.Printf("Задача № %d\n", i)
		}
	}()
}
func main() {

	Semaftor()

	for i:=1;i<=15;i++{
		jobs<-i
	}
	close(jobs)
	wg.Wait()
	close(result)

	}
