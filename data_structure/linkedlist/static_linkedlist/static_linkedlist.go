package linkedlist

// MAXSIZE  静态链表长度定义
var MAXSIZE = 10

// Node  静态链表结点
type Node struct {
	data   string //数据域
	cursor int    //游标
}

// Linkedlist 静态链表
type Linkedlist []Node

// Init 初始化
func Init(L Linkedlist) []Node {
	if L == nil {
		L = make([]Node, MAXSIZE)
	}
	for i := 0; i < MAXSIZE-1; i++ {
		L[i].cursor = i + 1 //将数组分量链接在一起成备用链表
	}
	L[MAXSIZE-2].cursor = 0 //备用链表的最后一个空元素的cur指向0
	L[MAXSIZE-1].cursor = 0 //链表最后一个结点的游标为0
	return L
}

// malloc  分配结点
func (L Linkedlist) malloc() (int, bool) {
	index := L[0].cursor
	if index == 0 {
		return -1, false
	}
	L[0].cursor = L[index].cursor
	return index, true
}

// free 将下标为k的空闲节点回收到备用链表中
func (L Linkedlist) free(index int) bool {
	if index < 1 || index > MAXSIZE {
		return false
	}
	L[index].cursor = L[0].cursor
	L[0].cursor = index
	return true
}

// Insert 插入结点
func (L Linkedlist) Insert(data string, index int) bool {
	if index < 1 || index > MAXSIZE {
		return false
	}

	return true
}
