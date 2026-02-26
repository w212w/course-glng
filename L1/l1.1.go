package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender string
}

func (h Human) Say() {
	fmt.Printf("Hello, my name is %s, I am %d years old and I am %s.\n", h.Name, h.Age, h.Gender)
}

type Action struct {
	Human
	ActionType string
}

func (a Action) Info() {
	fmt.Printf("%s has actionType: %s\n", a.Name, a.ActionType)
}

func main() {
	h := Human{Name: "Alice", Age: 30, Gender: "female"}
	a := Action{Human: h, ActionType: "running"}

	h.Say()
	a.Say()
	a.Info()
}
