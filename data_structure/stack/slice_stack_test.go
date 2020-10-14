package stack

import (
	"fmt"
	"testing"
)

func TestSliceStack(t *testing.T) {
	stack := make([]int, 0)
	Push(&stack, 5)
	Push(&stack, 10)

	peek := PeekS(&stack)
	v := Pop(&stack)
	fmt.Println(peek, v, stack)
}

func TestSliceStack1(t *testing.T) {
	s := make(stack, 0)
	s = s.Push(5)
	s = s.Push(10)

	peek := s.Peek()
	s, v := s.Pop()
	fmt.Println(peek, v, s)
}
