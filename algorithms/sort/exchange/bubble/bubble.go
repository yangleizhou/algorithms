package exchange

func bubble(nums []int) []int {
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func recursiveBubble(nums []int, size int) {
	if size == 1 {
		return
	}
	for i := 0; i < size-1; i++ {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
	recursiveBubble(nums, size-1)
}
