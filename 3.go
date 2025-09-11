package main

import (
	"fmt"
	"sync"
	"time"
)
func print(num<-chan int){ //BETA
	for req:=range filtered
}
func Semaftor(requests int, filtered chan<- int){
	for req := range requests{
		<-tick
		filtered<-req
	}
	close(filtered)
}
func main() {
	requests := make(chan int,15)
	ch2 := make(chan int,15)
	tick:= time.Tick(200*time.Millisecond)
}
