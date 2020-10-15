package stack

import (
	"fmt"
	"testing"
)

func TestSortStack(t *testing.T) {
	IStack := NewLinkedStack()
	var v = []int{5, 4, 9, 5, 8, 7, 7, 48}
	fmt.Println(v)
	IStack.Pushs(v)
	sortStack(IStack, compareIntGt)
	fmt.Println(IStack.Pops())
}
