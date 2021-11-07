package worker_pools

import (
	"fmt"
)

func WorkerPools() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 1. open activity monitor
	// 2. from concurrency_playground root directory, run `go run .`
	// 3. check activity. it should get close to 98%. kill the script
	// 4. now add three more workers.
	// 5. check activity 😎
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n - 1) + fib(n - 2)
}