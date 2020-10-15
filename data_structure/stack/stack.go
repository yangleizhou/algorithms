package stack

// https://blog.51cto.com/yinsuifeng/2359268
// https://blog.csdn.net/weixin_43258870/article/details/91353720
// https://www.sohu.com/a/270671146_478315

// IStack 栈方法接口
type IStack interface {
	Push(value interface{}) //元素入栈
	Pop() interface{}       //元素出栈
	Peek() interface{}      //查看栈顶元素
	IsEmpty() bool          //是否是空栈
	Size() int              //栈大小
}

type compFunc func(v1, v2 interface{}) bool

// sortStack 栈排序
func sortStack(s IStack, comp compFunc) {
	sortedStack := NewLinkedStack()
	var curElement interface{}
	for !s.IsEmpty() {
		curElement = s.Pop()
		for !sortedStack.IsEmpty() && comp(curElement, sortedStack.Peek()) {
			s.Push(sortedStack.Pop())
		}
		sortedStack.Push(curElement)
	}
	for !sortedStack.IsEmpty() {
		s.Push(sortedStack.Pop())
	}
}

func compareInt(v1, v2 interface{}) bool {
	return v1.(int) > v2.(int)
}
