package stack

// https://blog.51cto.com/yinsuifeng/2359268
// https://blog.csdn.net/weixin_43258870/article/details/91353720
// https://www.sohu.com/a/270671146_478315

// IStack 栈方法接口
type IStack interface {
	Push(value interface{}) //元素入栈
	Pop() interface{}       //元素出栈
	Peek() interface{}      //查看栈顶元素
	IsEmpty()               //是否是空栈
	Size() int              //栈大小
}
