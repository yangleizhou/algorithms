package stack

//通过slice 实现顺序栈

//Push  一个新项添加到栈的顶部
func Push(s *[]int, v int) {
	*s = append(*s, v)
}

//Pop 从栈中删除顶部项。它不需要参数并返回item。栈被修改
func Pop(s *[]int) int {
	if len(*s) == 0 {
		return 0
	}
	len := len(*s)
	v := (*s)[len-1]
	*s = (*s)[:len-1]
	return v
}

//Size 返回栈中的item数量。返回一个整数
func Size(s *[]int) int {
	return len(*s)
}

//IsEmpty 栈是否为空并返回布尔值
func IsEmpty(s *[]int) bool {
	return len(*s) == 0
}

//PeekS 栈返回顶部项，但不会删除它。不修改栈
func PeekS(s *[]int) int {
	if IsEmpty(s) {
		return 0
	}
	return (*s)[len(*s)-1]
}

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	len := len(s)
	if len == 0 {
		return s, 0
	}
	return s[:len-1], s[len-1]
}

func (s stack) IsEmpty() bool {
	return len(s) == 0
}

func (s stack) Peek() int {
	if s.IsEmpty() {
		return 0
	}
	return s[len(s)-1]
}
