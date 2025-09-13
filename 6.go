package main

import (
	"fmt"
	"time"
)

type Database struct {
	Name  string
	Speed time.Duration
}

func Search(db Database, ch chan<- string) {
	time.Sleep(db.Speed) //задержка в ответе от сайта
	result := db.Name
	ch <- result
}
func main() {
	dbs := []Database{
		{Name: "Google", Speed: 101 * time.Millisecond}, //искуственные базы для упрощения взаимодействия
		{Name: "Yandex", Speed: 500 * time.Millisecond},
		{Name: "VK MAX", Speed: 1 * time.Hour},
		{Name: "DuckDuckGo", Speed: 100 * time.Millisecond},
	}
	ch := make(chan string, 1)
	for _, ds := range dbs {
		go func(source Database) {
			Search(source, ch)
		}(ds)
	}
	fmt.Printf("Получен ответ от %v", <-ch)
}
