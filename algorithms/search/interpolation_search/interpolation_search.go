package search

// interpolationSearch 插值查找算法
func interpolationSearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	mid := low + (high-low)*(target-nums[low])/(nums[high]-nums[low])
	for low <= high {
		if target < nums[low] || target > nums[high] { // mid 范围可能不在[0,len(nums)-1]
			return -1
		}
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		}
		mid = low + (high-low)*(target-nums[low])/(nums[high]-nums[low])
	}
	return -1
}

// recursiveInterpolationSearch 递归实现插值查找算法
func recursiveInterpolationSearch(nums []int, low, high, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if low > high {
		return -1
	}
	if target < nums[low] || target > nums[high] { // mid 范围可能不在[0,len(nums)-1]
		return -1
	}
	mid := low + (high-low)*(target-nums[low])/(nums[high]-nums[low])
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return recursiveInterpolationSearch(nums, mid+1, high, target)
	} else {
		return recursiveInterpolationSearch(nums, low, high-1, target)
	}
}
