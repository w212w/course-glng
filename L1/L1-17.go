package main

import "fmt"

func binarySearch(a []int, target int) int {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)/2
		if a[mid] == target {
			return mid
		} else if a[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 4, 4, 6, 7, 8, 9, 12, 14, 15, 16, 18}
	target1 := 14
	target2 := 18

	fmt.Println(len(arr))
	fmt.Printf("index of %d: %d\n", target1, binarySearch(arr, target1))
	fmt.Printf("index of %d: %d\n", target2, binarySearch(arr, target2))
}
