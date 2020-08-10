package search

const size = 20 //斐波那契数组的长度

// fibonacciSearch 斐波那契搜索
func fibonacciSearch(nums []int, target int) int {
	n := len(nums)
	//
	k := 0
}

// fibonacci 构建斐波那契数组
func fibonacci(n int) []int {
	res := []int{0, 1}
	for i := 2; i < n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res
}
