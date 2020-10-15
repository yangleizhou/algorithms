package stack

// 最小栈

// MinStack 最小栈
type MinStack struct {
	Data *LinkedStack
	Min  *LinkedStack
}

// NewMinStack 获取最小栈空对象
func NewMinStack() *MinStack {
	return &MinStack{
		Data: &LinkedStack{},
		Min:  &LinkedStack{},
	}
}

// Push 入栈
func (s *MinStack) Push(v interface{}, comps ...compFunc) {
	s.Data.Push(v)
	var comp compFunc
	if len(comps) != 0 {
		comp = comps[0]
	} else {
		comp = compareIntLtorEt
	}
	if s.Min.IsEmpty() || comp(v, s.Min.Peek()) {
		s.Min.Push(v)
	}
}

// Pop 出栈
func (s *MinStack) Pop() interface{} {
	v := s.Data.Pop()
	if v == s.Min.Peek() {
		s.Min.Pop()
	}
	return v
}

// Peek 查看栈顶元素
func (s *MinStack) Peek() interface{} {
	return s.Data.Peek()
}

// GetMin 获取最小值
func (s *MinStack) GetMin() interface{} {
	return s.Min.Peek()
}
