package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	var N int
	fmt.Print("Enter timeout in seconds:")
	fmt.Scan(&N)
	timeout := time.Duration(N) * time.Second

	go func() {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}()

	timer := time.After(timeout)

	for {
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
		case <-timer:
			fmt.Println("Timeout reached, exiting...")
			return
		}
	}
}
