package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker1(id int, jobs <-chan int, done <-chan struct{}) {
	for {
		select {
		case j := <-jobs:
			fmt.Printf("Worker: %d processing job: %d\n", id, j)
		case <-done:
			fmt.Printf("Worker: %d stopping\n", id)
			return
		}
	}
}

func main() {
	var N int
	fmt.Print("Enter number of workers: ")
	fmt.Scan(&N)
	jobs := make(chan int)
	done := make(chan struct{})

	for w := 1; w <= N; w++ {
		go worker1(w, jobs, done)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		i := 1
		for {
			select {
			case <-done:
				close(jobs)
				return
			default:
				jobs <- i
				i++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-sigChan
	close(done)
	fmt.Println("Shutting down gracefully...")
	time.Sleep(2 * time.Second)
	fmt.Println("Program completed")

}
