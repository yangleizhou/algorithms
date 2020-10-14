package stack

//链式栈

//单向链表实现栈
type (
	//LinkedStack 链式栈结构体
	LinkedStack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// NewLinkedStack 获取链表栈对象
func NewLinkedStack() *LinkedStack {
	return &LinkedStack{}
}

//Size 返回栈大小
func (s *LinkedStack) Size() int {
	return s.length
}

//IsEmpty 是否是空栈
func (s *LinkedStack) IsEmpty() bool {
	return s.length == 0
}

//Peek  查看栈顶元素
func (s *LinkedStack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

//Pop 出栈
func (s *LinkedStack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push 入栈
func (s *LinkedStack) Push(v interface{}) {
	n := &node{v, s.top}
	s.top = n
	s.length++
}
