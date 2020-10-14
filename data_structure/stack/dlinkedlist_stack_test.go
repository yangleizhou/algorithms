package stack

import (
	"fmt"
	"testing"
)

func TestDLinkedlistStack(t *testing.T) {
	IStack := NewDLinkedStack()
	IStack.Push(5)
	IStack.Push(10)

	peek := IStack.Peek()
	v := IStack.Pop()
	fmt.Println(peek, v, IStack.Peek())
}
