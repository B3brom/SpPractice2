package main

import (
	"fmt"
	"net/http"
	"sync"
)

const PoolSize = 3

type Task string

type Result struct {
	URL        string
	StatusCode int
	Err        error
}

func worker(id int, jobs <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		resp, err := http.Get(string(job))
		if err != nil {
			results <- Result{URL: string(job), StatusCode: 0, Err: fmt.Errorf("ошибка: %w", err)}
			continue
		}
		results <- Result{URL: string(job), StatusCode: resp.StatusCode}
	}
}

func main() {
	urlsToScan := []string{
		"https://google.com",
		"https://yandex.ru",
		"https://nyarchlinux.moe",
		"https://discord.ru/",
		"https://несуществующийсайт.рф",
	}
	jobs := make(chan Task, len(urlsToScan))
	results := make(chan Result, len(urlsToScan))
	var wg sync.WaitGroup
	for i := 0; i < PoolSize; i++ {
		wg.Add(1)
		go worker(i+1, jobs, results, &wg)
	}
	for _, url := range urlsToScan {
		jobs <- Task(url)
	}
	close(jobs)
	wg.Wait()
	close(results)
	for res := range results {
		if res.Err != nil {
			fmt.Printf("URL: %s, Ошибка: %v\n", res.URL, res.Err)
		} else {
			fmt.Printf("URL: %s, Статус: %d\n", res.URL, res.StatusCode)
		}
	}
}
