package linkedlist

// MAXSIZE  静态链表长度定义
var MAXSIZE = 10

// Node  静态链表结点
type Node struct {
	data   int //数据域
	cursor int //游标
}

// Linkedlist 静态链表
type Linkedlist []Node

// Init 初始化
func Init(L Linkedlist) {
	for i := 0; i < MAXSIZE-1; i++ {
		L[i].cursor = i + 1 //将数组分量链接在一起成备用链表
	}
	L[MAXSIZE-1].cursor = 0 //链表最后一个结点的游标为0
}
