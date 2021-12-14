package oop

import "fmt"

type Huaman struct {
	Name string
	Age  int
}

func (h Huaman) Eat() {
	fmt.Println("human is eating...")
	h.Drink()
}

func (h Huaman) Drink() {
	fmt.Println("human is drinking...")
}

type Student struct {
	Huaman
	Grade int
	Major string
}

func (s Student) Drink() {
	fmt.Println("student is drinking...")
}
