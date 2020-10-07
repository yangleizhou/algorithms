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

// BinaryInsertion 折半插入排序
func BinaryInsertion(nums []int, p, q int) {
	for i := p + 1; i < q+1; i++ {
		low, high := p, i-1
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
}
