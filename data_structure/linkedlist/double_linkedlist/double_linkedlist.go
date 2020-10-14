package dlinkedlist

import "fmt"

// Node 链表结点
type Node struct {
	data int
	pre  *Node
	next *Node
}

// NewNode 获取新结点
func NewNode(v int) *Node {
	return &Node{data: v}
}

//GetData 获取数值
func (n *Node) GetData() int {
	return n.data
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
	l.head = &Node{}
	l.tail = &Node{}
	l.head.next = l.tail
	l.tail.pre = l.head
}

// Add 头部插入
func (l *LinkedList) Add(node *Node) bool {
	preNode := l.head
	node.next = preNode.next
	preNode.next = node

	node.pre = l.head
	node.next.pre = node
	l.size++
	return true
}

// Append 追加结点 (向后插入)
func (l *LinkedList) Append(node *Node) bool {
	if node == nil {
		return false
	}
	node.next = l.tail
	l.tail.pre.next = node

	node.pre = l.tail.pre
	l.tail.pre = node
	l.size++
	return true
}

// Insert 指定位置插入
// 0 插入首点结点 非0 插入对应值之后
func (l *LinkedList) Insert(i uint64, node *Node) bool {
	if node == nil {
		return false
	}
	if i > l.size {
		return false
	}
	preNode := l.head
	j := uint64(0)
	for j < i {
		preNode = preNode.next
		j++
	}
	node.next = preNode.next
	preNode.next = node

	node.pre = preNode
	node.next.pre = node
	l.size++
	return true
}

// Remove 移除指定位置的结点
func (l *LinkedList) Remove(i uint64) bool {
	if i >= l.size {
		return false
	}
	preNode := l.head
	j := uint64(0)
	for j < i {
		if preNode.next != nil {
			preNode = preNode.next
			j++
		}
	}
	node := preNode.next //找到删除结点
	preNode.next = node.next
	node.next.pre = preNode
	l.size--
	return true
}

// RemoveNode 移除指定结点
func (l *LinkedList) RemoveNode(node *Node) bool {
	if node == nil {
		return false
	}
	if l.size == 0 {
		return false
	}
	node.pre.next = node.next
	node.next.pre = node.pre
	l.size--
	return true
}

// RemoveAll 移除所有结点
func (l *LinkedList) RemoveAll() bool {
	l.Init()
	return true
}

// Get 获取指定位置的结点
func (l *LinkedList) Get(i uint64) *Node {
	if i >= l.size {
		return nil
	}
	node := l.head.next
	for j := uint64(0); j < i; j++ {
		node = node.next
	}
	return node
}

// GetHead 获取链表头
func (l *LinkedList) GetHead() *Node {
	return l.head.next
}

// GetTail 获取链表尾
func (l *LinkedList) GetTail() *Node {
	return l.tail.pre
}

// IndexOf 搜索某个数据的节点位置
func (l *LinkedList) IndexOf(data int) int {
	pos := -1
	if l.size == 0 {
		return pos
	}
	node := l.head.next
	for j := uint64(0); j < l.size; j++ {
		if node.data == data {
			pos = int(j)
			break
		}
		node = node.next
	}
	return pos
}

// GetLength 获取链表长度
func (l *LinkedList) GetLength() uint64 {
	return l.size
}

// FrontPrint 向后遍历
func (l *LinkedList) FrontPrint() {
	for head := l.head.next; head != l.tail; head = head.next {
		fmt.Printf("%d ", head.data)
	}
	fmt.Println()
}

// BackPrint 向前遍历
func (l *LinkedList) BackPrint() {
	if l.size == 0 {
		return
	}
	head := l.head
	for node := l.tail.pre; node != head; node = node.pre {
		fmt.Printf("%d ", node.data)
	}
	fmt.Println()
}
