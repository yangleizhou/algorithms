package stack

import (
	dlinkedlist "github.com/yangleizhou/algorithms/data_structure/linkedlist/double_linkedlist"
)

//链式栈

//双向链表实现

// DLinkedStack 双向链表链式栈结构体
type DLinkedStack struct {
	*dlinkedlist.LinkedList
}

// NewDLinkedStack 获取链表栈对象
func NewDLinkedStack() *DLinkedStack {
	l := &dlinkedlist.LinkedList{}
	l.Init()
	return &DLinkedStack{l}
}

//Size 返回栈大小
func (s *DLinkedStack) Size() int {
	return int(s.GetLength())
}

//IsEmpty 是否是空栈
func (s *DLinkedStack) IsEmpty() bool {
	return s.Size() == 0
}

//Peek  查看栈顶元素
func (s *DLinkedStack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	tail := s.GetTail()
	return tail.GetData()
}

//Pop 出栈
func (s *DLinkedStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	tail := s.GetTail()
	s.RemoveNode(tail)
	return tail.GetData()
}

//Push 入栈
func (s *DLinkedStack) Push(v interface{}) {
	node := dlinkedlist.NewNode(v.(int))
	s.Append(node)
}
