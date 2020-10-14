package stack

import (
	dclinkedlist "github.com/yangleizhou/algorithms/data_structure/linkedlist/circular_linkedlist/double_circular_linkedlist"
)

//链式栈

//双向循环链表实现

// DCLinkedStack 双向循环链表链式栈结构体
type DCLinkedStack struct {
	*dclinkedlist.LinkedList
}

// NewDCLinkedStack 获取链表栈对象
func NewDCLinkedStack() *DCLinkedStack {
	l := &dclinkedlist.LinkedList{}
	l.Init()
	return &DCLinkedStack{l}
}

//Size 返回栈大小
func (s *DCLinkedStack) Size() int {
	return int(s.GetLength())
}

//IsEmpty 是否是空栈
func (s *DCLinkedStack) IsEmpty() bool {
	return s.Size() == 0
}

//Peek  查看栈顶元素
func (s *DCLinkedStack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	tail := s.GetTail()
	return tail.GetData()
}

//Pop 出栈
func (s *DCLinkedStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	tail := s.GetTail()
	s.RemoveNode(tail)
	return tail.GetData()
}

//Push 入栈
func (s *DCLinkedStack) Push(v interface{}) {
	node := dclinkedlist.NewNode(v.(int))
	s.Append(node)
}
