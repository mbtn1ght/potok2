package main

import (
	"fmt"
	"sync"
)

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range jobs {
		results <- num * num
	}
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	jobs := make(chan int, len(slice))
	results := make(chan int, len(slice))

	for _, job := range slice {
		jobs <- job
	}
	close(jobs)

	var wg sync.WaitGroup
	numWorkers := 3
	for i := 0; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	wg.Wait()
	close(results)

	for _, num := range slice {
		fmt.Printf("Квадрат %v = %v\n", num, <-results)
	}
}
