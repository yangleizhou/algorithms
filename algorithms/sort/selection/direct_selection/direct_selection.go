package selection

func directSelection(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		// 寻找[i,n)区间里的最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		swap(nums, i, minIndex)
	}
	return nums
}

// 交换两数
// swap无法使用a=a+b,b=a-b,a=a-b;a=a^b,b=a^b,=a^b) 外层调用没有判断相等
func swap(nums []int, low, high int) {
	nums[low], nums[high] = nums[high], nums[low]
}
