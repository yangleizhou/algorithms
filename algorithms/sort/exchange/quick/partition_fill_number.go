package exchange

//双向扫描

// 挖坑填数 划分
func fillNumberPartition(nums []int, low, high, pivotType int) int {
	pivot := getPivot(nums, low, high, pivotType)
	i, j := low, high //左索引,右索引
	for i < j {
		// 由于j指向坑(nums[high])，所以需要先移动i
		// 而且每次需要检测i是否小于j，如果小于则继续
		// 否则退出移动，这样保证，退出for循环时，i=j,并且指向的是坑
		// 如果枢纽元在存放在数组第一个位置,此时i指向坑，流程只需要先移动j，其它不变
		for i < j && nums[i] < pivot { //实时监控 i< j 这里检测的是< pivot，等于就终止，为了保证元素都一样时，时间复杂度为O(NlogN)
			i++
		}
		if i < j { // 如果满足条件，进行填坑操作
			nums[j] = nums[i] // 将nums[i]填入nums[j]坑中，这时nums[i]变为了坑
			j--               // 同时对i自减，如果不这样，并且之前nums[i]=pivot的话，会造成死循环
		}
		for i < j && nums[j] > pivot {
			j--
		}
		if i < j {
			nums[i] = nums[j]
			i++
		}
	}
	nums[i] = pivot // 退出了循环，这是i=j，并且指向坑，将枢纽元填入坑中
	return i
}
