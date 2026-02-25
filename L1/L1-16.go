package main

import "fmt"

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	pivot := a[len(a)/2]
	var less, equal, greater []int
	for _, v := range a {
		if v < pivot {
			less = append(less, v)
		} else if v == pivot {
			equal = append(equal, v)
		} else {
			greater = append(greater, v)
		}
	}
	less = quickSort(less)
	greater = quickSort(greater)
	res := append(append(less, equal...), greater...)
	return res
}

func main() {
	arr := []int{8, 4, 3, 1, 6, 5, 7, 11, 9, 2, 10}
	fmt.Println("before:", arr)
	sorted := quickSort(arr)
	fmt.Println("after: ", sorted)
}
