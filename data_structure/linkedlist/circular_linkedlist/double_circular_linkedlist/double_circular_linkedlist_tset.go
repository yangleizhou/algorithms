package linkedlist

import "testing"

func TestDoubleCircularLinkedlist(t *testing.T) {
	l := &LinkedList{}
	l.Init()
	node1 := &Node{data: 1}
	node2 := &Node{data: 2}
	node3 := &Node{data: 3}
	node4 := &Node{data: 4}
	l.Add(node1)
	l.Append(node2)
	l.Append(node3)

	l.FrontPrint()
	l.BackPrint()

	l.Add(node4)

	l.FrontPrint()
	l.BackPrint()

	l.Remove(3)

	l.FrontPrint()
	l.BackPrint()
}
