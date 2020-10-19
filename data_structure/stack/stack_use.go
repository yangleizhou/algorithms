package stack

import "fmt"

// DecimalToAny 十进制转任意
func DecimalToAny(src int, baseType int) {
	IStack := NewLinkedStack()
	for src != 0 {
		IStack.Push(src % baseType)
		src /= baseType
	}
	for !IStack.IsEmpty() {
		fmt.Printf("%d", IStack.Pop())
	}
}

// BalanceSymbol 平衡符号
func BalanceSymbol(str string)
