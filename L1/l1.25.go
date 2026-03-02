package main

import (
	"fmt"
	"time"
)

func Sleep1(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
	}
}

func Sleep2(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	timer := 2 * time.Second
	fmt.Println("...")
	Sleep1(timer)
	fmt.Printf("Прошло %v секунды\n", timer.Seconds())

	fmt.Println("...")
	Sleep2(timer)
	fmt.Printf("Прошло %v секунды\n", timer.Seconds())
}
