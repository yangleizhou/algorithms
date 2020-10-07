package exchange

//双向扫描

// hoare 划分
func hoarePartition(nums []int, low, high, pivotType int) int {
	//选择基准
	pivot := getPivot(nums, low, high, pivotType)
	i, j := low-1, high //如果枢纽元存在的是数组第一个元素位置 i,j := low,high+1
	for i < j {
		i++
		j--
		for nums[i] < pivot {
			i++
		}
		for nums[j] > pivot {
			if j == low {
				break
			}
			j--
		}
		if i < j {
			swap(nums, i, j)
		}
	}
	// 此时i指向的元素大于等于pivot，如果大于pivot，那么是最左边大于pivot的元素
	// 交换i元素和最后一个元素即枢纽元
	swap(nums, i, high) //如果枢纽元存放的是数组的第一个位置 swap(nums,i,low)
	return i
}

// hoare 划分
func hoarePartitionMedianOptimize(nums []int, low, high int) int {
	pivot := getPivot(nums, low, high, MedianThreeOptimizePivot) //选择基准
	i := low                                                     //左索引
	j := high - 1                                                //右索引
	for i < j {
		i++                   //nums[i++]不会越界 & pivot 一定大于等于nums[i++]   golang 不支持nums[i++]
		j--                   //nums[j--]不会越界 & nums[j--] 一定大于等于pivot
		for nums[i] < pivot { //i向右扫描,j向左扫描，直至i>=j
			i++
		}
		for nums[j] > pivot {
			j--
		}
		if i < j {
			swap(nums, i, j)
		}
	}
	// 此时i指向的元素大于等于pivot，如果大于pivot，那么是最左边大于pivot的元素
	// 交换i元素和最后一个元素即枢纽元
	swap(nums, i, high-1)
	return i
}
