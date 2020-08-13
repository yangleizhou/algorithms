package insertion

// twoPathInsertionBinary 2-路插入排序算法
func twoPathInsertionBinary(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	first, last, n := 0, 0, len(nums)
	var temp = make([]int, n)
	temp[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < temp[first] {
			first = (first - 1 + n) % n
			temp[first] = nums[i]
		} else if nums[i] >= temp[last] {
			last++
			temp[last] = nums[i]
		} else { //[low,high)
			low, high := first, last
			for low != high { //first 不一定小于last 用 !=
				diff := (high - low + n) % n //元素个数
				mid := (low + (diff / 2)) % n
				if temp[mid] > nums[i] { // [high,last]一定比nums[i]大，所以但low == high 推出循环是[first,low-1) <=nums[i]<[low=high,last]
					high = mid // 如果使用high = (mid-1+n) %n  只能保证(high,last] 所以但 low==high 推出循环时 nums[i]与temp[low]大小无法确认
				} else {
					low = (mid + 1) % n //[first,(low-1+n)%n]一定比nums[i]小)
				}
			}

			// 移动元素
			for k := last + 1; k != low; k = (k - 1 + n) % n {
				temp[k] = temp[(k-1+n)%n]
			}
			temp[low] = nums[i]
			last++
		}
	}
	for i := 0; i < n; i++ {
		nums[i] = temp[(i+first)%n]
	}
	return nums
}
