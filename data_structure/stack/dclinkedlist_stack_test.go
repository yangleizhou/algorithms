package stack

import (
	"fmt"
	"testing"
)

func TestDCLinkedlistStack(t *testing.T) {
	IStack := NewDCLinkedStack()
	IStack.Push(5)
	IStack.Push(10)

	peek := IStack.Peek()
	v := IStack.Pop()
	fmt.Println(peek, v, IStack.Peek())
}
