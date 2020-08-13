package insertion

func binaryInsertion(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		low, high := 0, i-1
		for low <= high {
			mid := low + ((high - low) >> 1)
			if nums[mid] > nums[i] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		temp := nums[i]
		for j := i; j > high+1; j-- {
			nums[j] = nums[j-1]
		}
		nums[high+1] = temp
	}
	return nums
}