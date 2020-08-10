package search

// https://www.geeksforgeeks.org/fibonacci-search/

// fibonacciSearch2 不依赖数组实现斐波那契搜索
func fibonacciSearch2(nums []int, target int) int {
	var fibMMm2, fibMMm1 = 0, 1
	var fibM = fibMMm2 + fibMMm1

	n := len(nums)
	for fibM < n {
		fibMMm2 = fibMMm1
		fibMMm1 = fibM
		fibM = fibMMm2 + fibMMm1
	}
	//从前面开始，它标记着以删除的范围
	var offset = -1
	for fibM > 1 {
		i := min(offset+fibMMm2, n-1)
		if nums[i] < target {
			fibM = fibMMm1
			fibMMm1 = fibMMm2
			fibMMm2 = fibM - fibMMm1
			offset = i
		} else if nums[i] > target {
			fibM = fibMMm2
			fibMMm1 = fibMMm1 - fibMMm2
			fibMMm2 = fibM - fibMMm1
		} else {
			return i
		}
	}
	if fibMMm1 == 1 && (nums[offset+1] == target) {
		return offset + 1
	}
	return -1
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
