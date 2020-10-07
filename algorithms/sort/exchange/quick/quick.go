package exchange

func quickSortOneWay(nums []int, low, high int) {
	if low < high {
		//pivotIndex := oneWayPartition(nums, low, high, FixFirstPivot)
		//pivotIndex := oneWayPartition(nums, low, high, FixLastPivot)
		//pivotIndex := oneWayPartition(nums, low, high, RandPivot)
		pivotIndex := oneWayPartition(nums, low, high, MedianThreePivot)
		// 递归调用，这里之前被选择枢纽元的元素也参与了分割，所以不会将这个元素独立出来
		// 递归处理的两部分序列元素之和等于n而非n-1
		quickSortOneWay(nums, low, pivotIndex)
		quickSortOneWay(nums, pivotIndex+1, high)
	}
}

// 移动枢纽元 快排
func quickSortMovePivot(nums []int, low, high int) {
	if low < high {
		//pivotIndex := movePovitPartition(nums, low, high, FixFirstPivot)
		//pivotIndex := movePovitPartition(nums, low, high, FixLastPivot)
		//pivotIndex := movePovitPartition(nums, low, high, RandPivot)
		pivotIndex := movePovitPartition(nums, low, high, MedianThreePivot)
		// 递归调用，这里之前被选择枢纽元的元素也参与了分割，所以不会将这个元素独立出来
		// 递归处理的两部分序列元素之和等于n而非n-1
		quickSortMovePivot(nums, low, pivotIndex)
		quickSortMovePivot(nums, pivotIndex+1, high)
	}
}

// 挖坑填数 快排
func quickSortFillNumber(nums []int, low, high int) {
	if low < high {
		//pivotIndex := fillNumberPartition(nums, low, high, FixFirstPivot)
		//pivotIndex := fillNumberPartition(nums, low, high, FixLastPivot)
		//pivotIndex := fillNumberPartition(nums, low, high, RandPivot)
		pivotIndex := fillNumberPartition(nums, low, high, MedianThreePivot)
		quickSortFillNumber(nums, low, pivotIndex-1)
		quickSortFillNumber(nums, pivotIndex+1, high)
	}
}

// hoare划分 快排
func quickSortHoare(nums []int, low, high int) {
	if low < high {
		//pivotIndex := hoarePartition(nums, low, high, FixFirstPivot)
		//pivotIndex := hoarePartition(nums, low, high, FixLastPivot)
		//pivotIndex := hoarePartition(nums, low, high, RandPivot)
		pivotIndex := hoarePartition(nums, low, high, MedianThreePivot)
		quickSortHoare(nums, low, pivotIndex-1)
		quickSortHoare(nums, pivotIndex+1, high)
	}
}

// hoare划分(优化三数取中) 快排
func quickSortMedianOptimizeHoare(nums []int, low, high int) {
	if low < high {
		pivotIndex := hoarePartitionMedianOptimize(nums, low, high)
		quickSortMedianOptimizeHoare(nums, low, pivotIndex-1)
		quickSortMedianOptimizeHoare(nums, pivotIndex+1, high)
	}
}

// 三路 快排
func quickSortThreeWay(nums []int, low, high int) {
	if low < high {
		lt, gt := threeWayPartition(nums, low, high, MedianThreePivot)
		quickSortMedianOptimizeHoare(nums, low, lt)
		quickSortMedianOptimizeHoare(nums, gt, high)
	}
}
