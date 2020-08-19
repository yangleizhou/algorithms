package linkedlist

import (
	"fmt"
	"testing"
)

func TestSingleLinkedlist(t *testing.T) {
	var l = &LinkedList{}
	l.Init()
	node1 := &Node{data: 1}
	l.Add(node1)
	l.Append(&Node{data: 2})

	fmt.Println(l.GetHead().data, l.GetTail().data, l.size)

	l.Insert(0, &Node{data: 5})
	l.Insert(1, &Node{data: 6})
	fmt.Println(l.GetHead().data, l.GetTail().data, l.size)

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

	// for head := l.GetHead().next; head != nil; head = head.next {
	// 	fmt.Printf("%v\n", head)
	// }
}

func TestSingleLinkedlistRemove(t *testing.T) {
	var l = &LinkedList{}
	l.Init()

	l.Insert(0, &Node{data: 5})
	node1 := &Node{data: 1}
	l.Append(node1)
	l.Append(&Node{data: 2})

	l.Insert(2, &Node{data: 3})

	for head := l.GetHead(); head != nil; head = head.next {
		fmt.Printf("%v\n", head)
	}

	node4 := &Node{data: 6}
	l.Insert(0, node4)
	fmt.Println(l.GetHead().data, l.GetTail().data, l.size)

	l.RemoveNode(node4)
	fmt.Println(l.GetHead().data, l.GetTail().data, l.size)

	l.Insert(0, node4)
	fmt.Println(l.GetHead().data, l.GetTail().data, l.size)

	l.RemoveNode2(node4)
	fmt.Println(l.GetHead().data, l.GetTail().data, l.size)
}
