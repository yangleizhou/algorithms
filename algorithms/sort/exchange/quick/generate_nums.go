package exchange

import (
	"math/rand"
	"time"
)

// GenerateAscNums 生成升序数组
func GenerateAscNums(size int) []int {
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i + 1
	}
	return nums
}

// GenerateDesNums 生成降序数组
func GenerateDesNums(size int) []int {
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = size - i
	}
	return nums
}

// GenerateRandNums 随机数组
func GenerateRandNums(size int) []int {
	nums := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		nums[i] = rand.Int() % size
	}
	return nums
}

// GenerateRepeatNums 重复数组
func GenerateRepeatNums(size int) []int {
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		if i < size/2-1 {
			nums[i] = 10
		} else {
			nums[i] = 5
		}
	}
	return nums
}
