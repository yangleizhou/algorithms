package search

// 寻找左边界的二分查找
func leftBoundBinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums) //开区间
	mid := low + ((high - low) >> 1)
	for low < high { //开区间
		if nums[mid] == target {
			high = mid
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid
		}
		mid = low + ((high - low) >> 1)
	}
	return leftBoundReturnCmp(nums, low, target)
}

func leftBoundReturnCmp(nums []int, low, target int) int {
	if low == len(nums) {
		return -1
	}
	if nums[low] == target {
		return low
	}
	return -1
}
