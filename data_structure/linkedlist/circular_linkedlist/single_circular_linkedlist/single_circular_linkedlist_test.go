package linkedlist

import (
	"fmt"
	"testing"
)

func TestSingleCircularLinkedlist(t *testing.T) {
	var l = &LinkedList{}
	l.Init()
	node1 := &Node{data: 1}
	l.Add(node1)
	l.Append(&Node{data: 2})

	fmt.Println(l.GetHead().data, l.GetHead2().data, l.GetTail().data, l.size)

	l.Insert(0, &Node{data: 5})
	l.Insert(1, &Node{data: 6})
	fmt.Println(l.GetHead().data, l.GetHead2().data, l.GetTail().data, l.size)

	node := l.head.next
	for i := uint64(0); i < l.size; i++ {
		fmt.Print(node.data, ",")
		node = node.next
	}
	fmt.Println()

	l.Remove(0)
	node = l.head.next
	for i := uint64(0); i < l.size; i++ {
		fmt.Print(node.data, ",")
		node = node.next
	}
	fmt.Println()
	fmt.Println(l.IndexOf(2))
	fmt.Println(l.Get(0))

}
