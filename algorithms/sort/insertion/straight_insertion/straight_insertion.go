package insertion

func straightInsertion(nums []int) []int {
	length := len(nums)
	var preIndex, current int
	for i := 1; i < length; i++ {
		preIndex = i - 1
		current = nums[i]
		for preIndex >= 0 && nums[preIndex] > current {
			nums[preIndex+1] = nums[preIndex]
			preIndex--
		}
		nums[preIndex+1] = current
	}
	return nums
}
