package linkedlist

// 带头结点双向链表

// Node 链表结点
type Node struct {
	data int   //数据域
	next *Node //指针域
}

// LinkedList 单向链表
type LinkedList struct {
	size uint64 //链表大小
	head *Node  //头部结点
	tail *Node  //尾部结点
}

// Init 初始化链表
func (l *LinkedList) Init() {
	l.size = 0
	l.head = &Node{}
	l.tail = nil
	l.head.next = l.tail
}

// Append 追加结点 (向后插入)
func (l *LinkedList) Append(node *Node) bool {
	if node == nil {
		return false
	}
	head := l.head
	for head.next != nil {
		head = head.next
	}
	head.next = node
	l.tail = node
	l.size++
	return true
}

// Add 头部插入
func (l *LinkedList) Add(node *Node) bool {
	preNode := l.head
	node.next = preNode.next
	preNode.next = node
	if l.size == 0 {
		l.tail = node
	}
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
	if l.size == i {
		l.tail = node
	}
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
	node := preNode.next
	preNode.next = node.next

	if l.size-1 == i {
		l.tail = preNode
	}
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
	count := uint64(0)
	preNode := l.head
	for preNode.next != node {
		preNode = preNode.next
		count++
	}
	if count == l.size { //链表内不存在node结点
		return false
	}
	preNode.next = node.next

	if count == l.size-1 {
		l.tail = preNode
	}
	l.size--
	return true
}

// RemoveNode2 移除指定结点
// RemoveNode 优化点： 不必每次都遍历所有的节点，找到前驱节点
// 将这个需要删除的节点的后驱节点的数据域拷贝过来，然后删除这个后驱节点
func (l *LinkedList) RemoveNode2(node *Node) bool {
	if node == nil {
		return false
	}
	if l.size == 0 {
		return false
	}
	if node.next == nil { //最后一个节点，它其后没有后驱节点，所以需要从头遍历，找到它的前置节点
		preNode := l.head
		count := uint64(0)
		for preNode.next != node {
			preNode = preNode.next
			count++
		}
		if count == l.size {
			return false
		}
		preNode.next = node.next
		l.tail = preNode
	} else { //对于除最后一个节点的外的其他位置节点，则使用覆盖后删除后置节点的方式实现删除
		node.data = node.next.data
		node.next = node.next.next
	}
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
	return l.tail
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
