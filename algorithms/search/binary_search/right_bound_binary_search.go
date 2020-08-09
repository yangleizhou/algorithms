package search

// rightBoundBinarySearch 寻找右边界的二分查找
func rightBoundBinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums) //开区间
	mid := low + ((high - low) >> 1)
	for low < high { //开区间
		if nums[mid] == target {
			low = mid + 1
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid
		}
		mid = low + ((high - low) >> 1)
	}
	return rightBoundReturnCmp(nums, high, target)
}

func rightBoundReturnCmp(nums []int, high, target int) int {
	if high == len(nums) {
		return -1
	}
	if nums[high-1] == target {
		return high - 1
	}
	return -1
}

// recursiveRightBoundBinarySearch 递归实现右边界的二分查找
func recursiveRightBoundBinarySearch(nums []int, low, high, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if low == high {
		return rightBoundReturnCmp(nums, high, target)
	}
	mid := low + ((high - low) >> 1)
	if nums[mid] > target {
		return recursiveRightBoundBinarySearch(nums, low, mid, target)
	}
	return recursiveRightBoundBinarySearch(nums, mid+1, high, target)
}
