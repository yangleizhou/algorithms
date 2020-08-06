package binary

import "fmt"

// 解法一
func rangeBitwiseAnd1(m int, n int) int {
	if m == 0 {
		return 0
	}
	moved := 0
	for m != n {
		if n == 0 {
			return 0
		}
		m >>= 1
		n >>= 1
		moved++
	}
	fmt.Println(moved)
	return m << uint32(moved)
}

// 解法二 Brian Kernighan's algorithm
func rangeBitwiseAnd2(m int, n int) int {
	var count int
	for n > m {
		count++
		n &= (n - 1) // 清除最低位的 1
	}
	fmt.Println(count)
	return n
}
