package stack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(5)
	minStack.Push(5)
	minStack.Push(8)
	minStack.Push(7)
	minStack.Push(6)
	minStack.Push(3)
	fmt.Println(minStack.GetMin())
	fmt.Println(minStack.Pop(), minStack.GetMin())
	fmt.Println(minStack.Pop(), minStack.GetMin())
	fmt.Println(minStack.Pop(), minStack.GetMin())
	fmt.Println(minStack.Pop(), minStack.GetMin())
	fmt.Println(minStack.Pop(), minStack.GetMin())
}

// 3
// 3 5
// 6 5
// 7 5
// 8 5
// 5 5

func TestMinStack1(t *testing.T) {
	minStack := NewLinkedStack()
	minStack.PushMinStack(5)
	minStack.PushMinStack(5)
	minStack.PushMinStack(8)
	minStack.PushMinStack(7)
	minStack.PushMinStack(6)
	minStack.PushMinStack(3)
	fmt.Println(minStack.GetMin())
	fmt.Println(minStack.Pop(), minStack.GetMin())
	fmt.Println(minStack.Pop(), minStack.GetMin())
	fmt.Println(minStack.Pop(), minStack.GetMin())
	fmt.Println(minStack.Pop(), minStack.GetMin())
	fmt.Println(minStack.Pop(), minStack.GetMin())
}
