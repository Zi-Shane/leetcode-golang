package stack

import (
	"sync"
)

type (
	Stack struct {
		top    *node
		length int
		lock   *sync.RWMutex
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func NewStack() *Stack {
	return &Stack{nil, 0, &sync.RWMutex{}}
}

// Return the number of item in the stack
func (stack *Stack) Len() int {
	return stack.length
}

// Return true if stack is empty
func (stack *Stack) IsEmpty() bool {
	return stack.length == 0
}

func (stack *Stack) PushInt(value int) {
	stack.lock.Lock()
	stack.top = &node{value: value, prev: stack.top}
	stack.length++
	stack.lock.Unlock()
}

func (stack *Stack) PopInt() int {
	stack.lock.Lock()
	if stack.length == 0 {
		return -1
	}
	item := stack.top
	stack.top = item.prev
	stack.length--
	stack.lock.Unlock()
	return item.value.(int)
}
