package stack

import (
	"fmt"
	"testing"
)

func TestLinkedlistStack(t *testing.T) {
	IStack := NewLinkedStack()
	IStack.Push(5)
	IStack.Push(10)
	IStack.Push(20)

	peek := IStack.Peek()
	v := IStack.Pop()
	fmt.Println(peek, v, IStack.Peek(), IStack.Pop())
}
