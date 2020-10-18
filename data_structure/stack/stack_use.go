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
func BalanceSymbol(s IStack, chars []rune) bool {
	if s == nil || len(chars) == 0 {
		return false
	}
	for index := range chars {
		c := chars[index]
		switch c {
		case '[', '{', '(':
			s.Push(c)
		case ']', '}', ')':
			v := s.Pop()
			if v == nil {
				return false
			}
			return balanceSymbolJudge(c, v.(rune))
		default:
		}
	}
	return true
}

func balanceSymbolJudge(char, src rune) bool {
	switch char {
	case '(':
		return src == ')'
	case '}':
		return src == '{'
	case ']':
		return src == '['
	}
	return false
}
