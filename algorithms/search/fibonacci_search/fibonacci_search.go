package search

// fibonacciSearch 斐波那契搜索
func fibonacciSearch(nums []int, target int) int {
	n := len(nums)
	fib, k := fibonacci(n)
	filled := getFilled(nums, fib[k]-1)
	low, high := 0, n-1
	mid := low + fib[k-1] - 1
	for low <= high {
		if filled[mid] == target {
			if mid >= n { // 在填充数组中查找成功
				mid = n - 1
			}
			return mid
		} else if filled[mid] < target {
			low = mid + 1
			k -= 2
		} else if filled[mid] > target {
			high = mid - 1
			k--
		}
		mid = low + fib[k-1] - 1
	}
	return -1
}

// getFilled 将原数组查找表扩展为长度为F[k]-1的填充数组
func getFilled(nums []int, k int) (filled []int) {
	n := len(nums)
	filled = make([]int, k)
	copy(filled, nums)
	for i := n; i < k; i++ {
		filled[i] = nums[n-1]
	}
	return
}

// fibonacci 构建斐波那契数组
// 根据待查找数组长度确定裴波那契数组 满足F(k) - 1 >=n
func fibonacci(n int) ([]int, int) {
	res := []int{0, 1}
	k := 0
	for k = 2; ; k++ {
		res = append(res, res[k-1]+res[k-2])
		if res[k]-1 >= n {
			break
		}
	}
	return res, k
}
