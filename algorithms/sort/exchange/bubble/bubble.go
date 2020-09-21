package exchange

func bubble(nums []int) []int {
	size := len(nums)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size-1; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func bubbleSort(nums []int, size int) {
	for i := 0; i < size-1; i++ {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
	bubbleSort(nums, size-1)
}
