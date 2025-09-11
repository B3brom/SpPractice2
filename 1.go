package main

import (
	"fmt"
	"sync"
	"time"
)
func print(wg *sync.WaitGroup){
	defer wg.Done() 
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1) 
	go print(&wg)
	wg.Wait()
}
