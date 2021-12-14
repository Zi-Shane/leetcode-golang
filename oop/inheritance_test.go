package oop_test

import (
	"leetcode/oop"
	"testing"
)

func TestInheritance(t *testing.T) {
	s := oop.Student{oop.Huaman{"Alice", 16}, 100, "CS"}
	s.Drink()
	s.Eat()
}
