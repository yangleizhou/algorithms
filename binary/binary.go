package binary

import "fmt"

// 判断一个数的奇偶性 `n & 1`
func parity(x int) {
	if (x & 1) == 0 {
		fmt.Printf("%d 是偶数\n", x)
	} else {
		fmt.Printf("%d 是奇数\n", x)
	}
}

// 利用或操作 | 和空格将英文字符转换为小写
func convertLower(strArr []byte) {
	for index, str := range strArr {
		strArr[index] = (str | ' ')
	}
}

// 利用与操作 `&` 和下划线将英文字符转换为大写
func convertSupper(strAttr []byte) {
	for index, str := range strAttr {
		strAttr[index] = (str & '_')
	}
}

// 利用异或操作 `^` 和空格进行英文字符大小写互换
func convertLowerSupper(strAttr []byte) {
	for index, str := range strAttr {
		strAttr[index] = (str ^ ' ')
	}
}

// 利用 `^` 判断两个数是否异号
func isSameSymbol(x, y, z int) {
	if (x ^ y) > 0 {
		fmt.Printf("x = %d,y =%d 相同符号\n", x, y)
	} else {
		fmt.Printf("x = %d,y =%d 不同符号\n", x, y)
	}

	if (x ^ z) > 0 {
		fmt.Printf("x = %d,z =%d 相同符号\n", x, z)
	} else {
		fmt.Printf("x = %d,z =%d 不同符号\n", x, z)
	}
}

// 利用 `^` 交换两个数
func swap(x, y int) {
	fmt.Println(x, y)
	x ^= y
	y ^= x
	x ^= y
	fmt.Println(x, y)
}

// 利用`n & (n-1)`判断一个整数是不是2的指数
func isPowerOf2(x int) bool {
	if x <= 0 {
		return false
	}
	return x&(x-1) == 0
}

// 位移实现 `>>` 或 n & (n-1) 计算汉明权重
func hammingWeight(x, y uint32) {
	// 使用 >>
	res := uint32(0)
	for x != 0 {
		res += (x & 0X01)
		x = x >> 1
	}
	fmt.Println(res)
	res = 0

	// 使用 n & (n-1)
	for y != 0 {
		y = y & (y - 1)
		res++
	}
	fmt.Println(res)
}
