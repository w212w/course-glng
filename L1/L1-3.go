package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int) {
	for j := range jobs {
		fmt.Printf("Worker: %d processing job: %d\n", id, j)
	}
}

func main() {
	var N int
	fmt.Print("Enter number of workers: ")
	fmt.Scan(&N)

	jobs := make(chan int)

	for w := 1; w <= N; w++ {
		go worker(w, jobs)
	}

	i := 1
	for {
		jobs <- i
		i++
		time.Sleep(500 * time.Millisecond)
	}
}
