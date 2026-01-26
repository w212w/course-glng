package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func stopByCondition() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("stopByCondition:", i)
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("stopByCondition: done")
	}()
	time.Sleep(2 * time.Second)
}

func stopByChannel() {
	stop := make(chan struct{})
	go func() {
		i := 0
		for {
			select {
			case <-stop:
				fmt.Println("stopByChannel: done")
				return
			default:
				fmt.Println("stopByChannel:", i)
				i++
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	stop <- struct{}{}
	time.Sleep(500 * time.Millisecond)
}

func stopByContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopByContext: done")
				return
			default:
				fmt.Println("stopByContext:", i)
				i++
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}

func stopByContextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopByContextWithTimeout: done")
				return
			default:
				fmt.Println("stopByContextWithTimeout:", i)
				i++
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(4 * time.Second)
}

func stopByContextWithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()
	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopByContextWithDeadline: done")
				return
			default:
				fmt.Println("stopByContextWithDeadline:", i)
				i++
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(4 * time.Second)
}

func stopByGoexit() {
	go func() {
		defer fmt.Println("stopByGoexit: deferred call executed")
		for i := 0; i < 3; i++ {
			fmt.Println("stopByGoexit:", i)
			time.Sleep(500 * time.Millisecond)
		}
		runtime.Goexit()
		fmt.Println("stopByGoexit: unreachable")
	}()
	time.Sleep(2 * time.Second)
}

func main() {
	stopByCondition()
	stopByChannel()
	stopByContext()
	stopByContextWithTimeout()
	stopByContextWithDeadline()
	stopByGoexit()
	time.Sleep(3 * time.Second)
}
