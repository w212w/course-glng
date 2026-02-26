package main

func main() {
	nums := []int{10, 20, 30, 40, 50}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, num := range nums {
			ch1 <- num
		}
		close(ch1)
	}()

	go func() {
		for num := range ch1 {
			ch2 <- num * 2
		}
		close(ch2)
	}()

	for num := range ch2 {
		println(num)
	}
}
