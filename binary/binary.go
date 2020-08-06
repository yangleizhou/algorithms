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

// n & (n-1) 利用 n & (n- 1) 移除最后一个1
func removeLast1(x int) int {
	return x & (x - 1)
}

// 利用 (n & (n- 1))^n 获取最后一个1
func getLast1(x int) int {
	return (x & (x - 1)) ^ x
}

// 利用`n & (n-1)`判断一个整数是不是2的指数
func isPowerOf2(x int) bool {
	if x <= 0 {
		return false
	}
	return x&(x-1) == 0
}

// 位移实现 `>>` 或 n & (n-1) 计算汉明权重
// n & (n-1) 去掉最后一位
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

// 利用 a ^ a = 0 ^ b = b
// 找到非空整型数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
func singleNumber(data ...int) int {
	var result int
	for _, v := range data {
		result = result ^ v
	}
	return result
}

// 统计各个数各位上1的数量 在%n
// 找到非空整型数组，除了某个元素只出现一次以外，其余每个元素均出现n次。找出那个只出现了一次的元素
func singleNumber2(n int, data []int) int {
	if n <= 1 {
		return 0
	}
	var result int
	for i := 0; i < 64; i++ {
		sumbit := 0
		for j := 0; j < len(data); j++ {
			sumbit += (data[j] >> i) & 1
		}
		result ^= (sumbit % n) << i
	}
	return result
}

// 利用 a^a=0^b=b 出现两次数 ^ = 0 最后f= a^b ;两数交换 f= a^b,a = f^b,b=f^a
// f 最后一个1=g,g&a=0 g&b !=0 或者 g&a !=0 g&b=0
// 给定一个整数数组，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素
func singleNumber3(nums []int) []int {
	// a=a^b
	// b=a^b
	// a=a^b
	// 关键点怎么把a^b分成两部分,方案：可以通过diff最后一个1区分

	diff := 0
	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}
	result := []int{diff, diff}
	// 去掉末尾的1后异或diff就得到最后一个1的位置
	diff = (diff & (diff - 1)) ^ diff
	for i := 0; i < len(nums); i++ {
		if diff&nums[i] == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	return result
}
