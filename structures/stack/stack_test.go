package Stack_test

import (
	"fmt"
	Stack "leetcode/structures/stack"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack.NewStack()
	s.PushInt(1)
	fmt.Println("stack length: ", s.Len())
	if !s.IsEmpty() {
		fmt.Println("do pop() got value: ", s.PopInt())
	}
	fmt.Println("stack length: ", s.Len())
}
