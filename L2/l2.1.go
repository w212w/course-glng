package main

import "fmt"

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}

// 77, 78, 79
// Срез b берёт элементы из a начиная с 1 включительно до 4 (не включительно).
// При этом len(b) = 3, cap(b) = 5-1 = 4
