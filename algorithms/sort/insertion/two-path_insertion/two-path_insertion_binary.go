package insertion

// 2-路插入排序算法 twoPathInsertionBinary
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
		} else {
			low, high := first, last
			for low <= high {
				diff := (high - low + n) % n
				mid := (low + diff>>1) % n
				if nums[mid] > nums[i] {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}
		}
	}
}
