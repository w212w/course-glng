package main

import "fmt"

func removeElement[T any](slice []T, i int) []T {
	if i < 0 || i >= len(slice) {
		panic("index out of range")
	}

	copy(slice[i:], slice[i+1:])

	var zero T
	slice[len(slice)-1] = zero
	return slice[:len(slice)-1]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("До удаления:", nums)
	nums = removeElement(nums, 2)
	fmt.Println("После удаления элемента с индексом 2:", nums)

	words := []string{"apple", "banana", "cherry", "date"}
	fmt.Println("\nДо удаления:", words)
	words = removeElement(words, 1)
	fmt.Println("После удаления элемента с индексом 1:", words)

	type User struct {
		Name string
		Age  int
	}

	users := []*User{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	fmt.Println("\nДо удаления:")

	for _, u := range users {
		fmt.Printf("  %+v\n", u)
	}

	users = removeElement(users, 1)

	fmt.Println("После удаления элемента с индексом 1:")

	for _, u := range users {
		fmt.Printf("  %+v\n", u)
	}
}
