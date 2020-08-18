package linkedlist

// Node 链表结点
type Node struct {
	data int
	next *Node
}

// LinkedList 单向循环链表
type LinkedList struct {
	size uint64 //链表大小
	head *Node  //头部结点
	tail *Node  //尾部结点
}

// Init 初始化链表
func (l *LinkedList) Init() {
	l.size = 0
	l.head = &Node{}
	l.tail.next = l.head
}
