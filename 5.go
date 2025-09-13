package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"sync"
)

const WorkerPoolSize = 3

func main() {
	files := []string{"1.txt", "2.txt", "3.txt", "4.txt", "5.txt", "6.txt", "7.txt"} //Code runner не видит файлы. Запускать только через F5
	semaphore := make(chan struct{}, WorkerPoolSize)                                 //UPD: Code Runner начинает работать если в настройках расширения включить "Code-runner: File Directory As Cwd"
	var wg sync.WaitGroup                                                            //Тогда среда компиляции(или выполнения) будет находиться в папке запускаемого файла
	var mu sync.Mutex

	fmt.Printf("Готов обработать %d файлов пачками по %d файлика.\n", len(files), WorkerPoolSize)

	for _, f := range files {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(filePath string) {
			defer wg.Done()
			defer func() { <-semaphore }()
			mu.Lock()
			fmt.Printf("Обработка: %s\n", filePath)
			mu.Unlock()
			data, err := os.ReadFile(filePath)
			if err != nil {
				mu.Lock()
				fmt.Printf("Чет с файлом не клеится %s: %v\n", filePath, err)
				mu.Unlock()
				return
			}
			hash := md5.Sum(data)
			mu.Lock()
			fmt.Printf("%s: %x\n", filePath, hash)
			mu.Unlock()
		}(f)
	}

	wg.Wait()
	fmt.Println("Дело сделано")
}
