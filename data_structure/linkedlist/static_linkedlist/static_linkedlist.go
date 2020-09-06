package linkedlist

import "fmt"

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
func Init() Linkedlist {
	L := make([]Node, MAXSIZE)
	for i := 0; i < MAXSIZE-1; i++ {
		L[i].cursor = i + 1 //将数组分量链接在一起成备用链表
	}
	L[MAXSIZE-2].cursor = 0 //备用链表的最后一个空元素的cur指向0
	L[MAXSIZE-1].cursor = 0 //链表最后一个结点的游标为0
	return L
}

// malloc  分配结点
func (L Linkedlist) malloc() (int, bool) {
	if L.isFull() {
		return -1, false
	}
	index := L[0].cursor
	L[0].cursor = L[index].cursor
	return index, true
}

// free 将下标为k的空闲节点回收到备用链表中
func (L Linkedlist) free(index int) bool {
	if index < 1 || index > MAXSIZE-1 {
		return false
	}
	L[index].cursor = L[0].cursor
	L[0].cursor = index
	return true
}

//回收链表到备用链表
func (L Linkedlist) destroy() {
	if L[MAXSIZE-1].cursor == 0 {
		return
	}
	i := L[MAXSIZE-1].cursor
	L[MAXSIZE-1].cursor = 0

	idelHeadIndex := L[0].cursor
	L[0].cursor = i

	lastIndex := i
	for i > 0 {
		lastIndex = i
		i = L[i].cursor
	}
	L[lastIndex].cursor = idelHeadIndex
}

// 静态链表是否为空
func (L Linkedlist) isEmpty() bool {
	if L[MAXSIZE-1].cursor == 0 {
		return true
	}
	return false
}

// 静态链表是否满
func (L Linkedlist) isFull() bool {
	return L[0].cursor == 0
}

// 静态链表长度
func (L Linkedlist) length() uint64 {
	i := L[MAXSIZE-1].cursor
	if i == 0 || i == L[0].cursor {
		return 0
	}
	j := 0
	for i > 0 {
		j++
		i = L[i].cursor
	}
	return uint64(j)
}

//遍历链表
func (L Linkedlist) trave() {
	for i := L[MAXSIZE-1].cursor; i > 0; {
		fmt.Println(i, ":", L[i].data)
		i = L[i].cursor
	}
}

// 获取index索引对应位置
func (L Linkedlist) getIndexPos(index int) int {
	i := L[MAXSIZE-1].cursor
	if i == 0 {
		return -1
	}
	count := 1
	for i > 0 && count < index-1 {
		count++
		i = L[i].cursor
	}
	return i
}

func (L Linkedlist) firstAdd(data string) bool {
	curIndex, ok := L.malloc()
	if !ok {
		return false
	}
	L[curIndex].data = data
	L[curIndex].cursor = 0
	if L[MAXSIZE-1].cursor == 0 {
		L[MAXSIZE-1].cursor = curIndex
	}
	return true
}

func (L Linkedlist) add(data string, index int) bool {
	i := L.getIndexPos(index)
	if i == -1 {
		return false
	}

	curIndex, ok := L.malloc()
	if !ok {
		return false
	}
	L[curIndex].data = data
	L[curIndex].cursor = L[i].cursor
	L[i].cursor = curIndex
	return true
}

// 尾插
func (L Linkedlist) Append(data string) bool {
	i := L[MAXSIZE-1].cursor
	if i == 0 {
		L.firstAdd(data)
		return true
	}
	lastIndex := i
	for i > 0 {
		lastIndex = i
		i = L[i].cursor
	}
	index, ok := L.malloc()
	if !ok {
		return false
	}
	L[lastIndex].cursor = index
	L[index].data = data
	L[index].cursor = 0
	return true
}

// 头插
func (L Linkedlist) HeadAdd(data string) bool {
	i := L[MAXSIZE-1].cursor
	if i == 0 {
		L.firstAdd(data)
		return true
	}
	index, ok := L.malloc()
	if !ok {
		return false
	}
	temp := L[MAXSIZE-1].cursor
	L[index].data = data
	L[index].cursor = temp
	L[MAXSIZE-1].cursor = index
	return true
}

// Insert 插入结点
func (L Linkedlist) Insert(data string, index int) bool {
	if index < 1 || (index-1) > int(L.length()) {
		return false
	}

	i := L[MAXSIZE-1].cursor
	if i == 0 {
		L.firstAdd(data)
		return true
	}
	return L.add(data, index)
}

// Remove 删除结点
func (L Linkedlist) Remove(index int) bool {
	if index < 1 || (index-1) > int(L.length()) {
		return false
	}
	i := L.getIndexPos(index)
	if i == -1 {
		return false
	}
	temp := L[i].cursor
	L[i].cursor = L[temp].cursor
	L.free(temp)
	return true
}
