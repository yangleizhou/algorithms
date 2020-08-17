package linkedlist

// Node 链表结点
type Node struct {
	data int
	pre  *Node
	next *Node
}

// LinkedList 双向链表
type LinkedList struct {
	size uint64
	head *Node //头结点
	tail *Node //尾结点
}

// Init 初始化链表
func (l *LinkedList) Init() {
	l.size = 0
	l.head = nil
	l.tail = nil
}
