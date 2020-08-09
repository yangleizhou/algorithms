package search

// binarySearch 二分查找
func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1 //闭区间
	mid := low + ((high - low) >> 1)
	for low <= high { //闭区间
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		}
		mid = low + ((high - low) >> 1)
	}
	return -1
}

// recursiveBinarySearch 递归算法实现二分查找
func recursiveBinarySearch(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := low + ((high - low) >> 1)

	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return recursiveBinarySearch(nums, mid+1, high, target)
	} else {
		return recursiveBinarySearch(nums, low, high-1, target)
	}
}
