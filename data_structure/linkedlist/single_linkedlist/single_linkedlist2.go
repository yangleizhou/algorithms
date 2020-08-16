package linkedlist

// 不带头结点单向链表

// LinkedList2 单向链表
type LinkedList2 struct {
	size uint64 //链表大小
	head *Node  //首元结点
	tail *Node  //尾部结点
}

// Init 初始化链表
func (l *LinkedList2) Init() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

// Append 追加结点 (向后插入)
func (l *LinkedList2) Append(node *Node) bool {
	if node == nil {
		return false
	}
	node.next = nil  //新追加结点在末尾,没有next
	if 0 == l.size { //是首元结点
		l.head = node //nil 链表，node 成为头结点
	} else {
		l.tail.next = node //尾部next 指向新结点
	}
	l.tail = node
	l.size++
	return true
}

// Insert 向第i个结点前插入结点 (向前插入)
func (l *LinkedList2) Insert(i uint64, node *Node) bool {
	if node == nil {
		return false
	}
	if i > l.size {
		return false
	}

	if 0 == i { //是首元结点
		node.next = l.head
		l.head = node
		if l.size == 0 {
			l.tail = node
		}
	} else {
		pre := l.head
		j := uint64(1)
		for j < i { //寻找前驱指针
			pre = pre.next
			j++
		}
		node.next = pre.next // 新结点指向旧结点原来所指向的next
		pre.next = node      // 原结点的next指向新结点
	}
	l.size++
	return true
}

// Remove 移除指定位置的结点
func (l *LinkedList2) Remove(i uint64) bool {
	if i >= l.size {
		return false
	}
	if 0 == i { //是首元结点
		l.head = l.head.next //旧头部的next变为新头部
	} else {
		preNode := l.head
		j := uint64(1)
		for j < i { //寻找前驱指针
			preNode = preNode.next
		}
		node := preNode.next     //找到当前要删除的结点
		preNode.next = node.next //把当前要删除结点的next赋值给其父节点的next
		if i == l.size-1 {
			l.tail = preNode
		}
	}
	l.size--
	return true
}

// RemoveNode 移除指定结点
func (l *LinkedList2) RemoveNode(node *Node) bool {
	if node == nil {
		return false
	}
	if l.size == 0 {
		return false
	}
	count := uint64(0)
	preNode := l.head
	for preNode != node {
		preNode = preNode.next
		count++
	}
	if count == l.size { //链表内不存在node结点
		return false
	}
	if count == 0 { //是首元结点
		l.head = node.next
	} else {
		preNode.next = node.next
		if count == l.size-1 {
			l.tail = preNode
		}
	}
	l.size--
	return true
}

// RemoveNode2 移除指定结点
// RemoveNode 优化点： 不必每次都遍历所有的节点，找到前驱节点
// 将这个需要删除的节点的后驱节点的数据域拷贝过来，然后删除这个后驱节点
func (l *LinkedList2) RemoveNode2(node *Node) bool {
	if node == nil {
		return false
	}
	if l.size == 0 {
		return false
	}
	if node.next == nil { //最后一个节点，它其后没有后驱节点，所以需要从头遍历，找到它的前置节点
		preNode := l.head
		count := uint64(0)
		for preNode != node {
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
func (l *LinkedList2) RemoveAll() bool {
	l.Init()
	return true
}

// Get 获取指定位置的结点
func (l *LinkedList2) Get(i uint64) *Node {
	if i >= l.size {
		return nil
	}
	node := l.head
	for j := uint64(0); j < i; j++ {
		node = node.next
	}
	return node
}

// GetHead 获取链表头
func (l *LinkedList2) GetHead() *Node {
	return l.head
}

// GetTail 获取链表尾
func (l *LinkedList2) GetTail() *Node {
	return l.tail
}

// IndexOf 搜索某个数据的节点位置
func (l *LinkedList2) IndexOf(data int) int {
	pos := -1
	if l.size == 0 {
		return pos
	}
	node := l.head
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
func (l *LinkedList2) GetLength() uint64 {
	return l.size
}
