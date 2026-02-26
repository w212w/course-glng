package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 3, 5, 6, 1, 3, 2}
	b := []int{2, 2, 2, 4, 5, 6, 7, 8}
	fmt.Println("Пересечение =", arrays(a, b))

}

func arrays(a, b []int) []int {
	mapa := make(map[int]struct{}, len(a))
	for _, v := range a {
		mapa[v] = struct{}{}
	}
	var result []int
	seen := make(map[int]struct{})
	for _, v := range b {
		if _, exists := mapa[v]; exists {
			if _, ok := seen[v]; !ok {
				result = append(result, v)
				seen[v] = struct{}{}
			}
		}

	}
	return result
}
