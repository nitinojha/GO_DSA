package main

import (
	"fmt"
)

type Stack struct {
	element  []int
	capacity int
	top      int
}

func NewStack(cap int) *Stack {
	return &Stack{element: make([]int, 0, cap), capacity: cap, top: 0}
}

func (s *Stack) Push(val int) bool {
	if len(s.element) >= s.capacity {
		fmt.Println("Stack is overflowed")
		return false
	} else {
		s.element = append(s.element, val)
		s.top++
		return true
	}
}

func (s *Stack) Pop() (int, bool) {
	if len(s.element) != 0 {
		popped_element := s.element[len(s.element)-1]
		s.element = s.element[:len(s.element)-1]
		fmt.Println("hryy", s.element, len(s.element)-1)
		s.top--
		return popped_element, true
	}
	return 0, false
}

func (s *Stack) IsEmpty() bool {
	return len(s.element) == 0
}

func (s *Stack) Top() (int, bool) {
	if s.top < 1 {
		return 0, false
	} else {
		return s.element[s.top-1], true
	}
}

func main() {
	stack := NewStack(5)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	fmt.Println("1", stack.element)
	current_top, _ := stack.Top()
	fmt.Println("1top", current_top)
	current_pop, _ := stack.Pop()
	fmt.Println("1Pop", current_pop)
	for i := 0; i < 3; i++ {
		if stack.IsEmpty() {
			stack.Push(50)
		} else {
			current_top, _ := stack.Top()
			fmt.Println("1top", current_top)
			current_pop, _ := stack.Pop()
			fmt.Println("1Pop", current_pop)
		}
		fmt.Println(i, stack.element)
	}
}
