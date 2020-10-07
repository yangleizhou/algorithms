package exchange

//3路划分

// lt 表示<pivot部分的最后一个元素所在位置
// gt 表示>pivot部分的第一个元素所在位置
func threeWayPartition(nums []int, low, high, pivotType int) (int, int) {
	//选择基准
	pivot := getPivot(nums, low, high, pivotType)
	i := low            //[如果枢纽元放在第一个位置 i:= low+1]
	lt, gt := low, high //[如果枢纽元放在第一个位置 lt,gt:= low,high+1]
	for i < gt {
		if nums[i] == pivot {
			i++
		} else if nums[i] > pivot {
			// 如果当前位置元素>pivot，则将当前位置元素与>pivot部分的第一个元素的前一个元素交换位置
			swap(nums, gt-1, i)
			//表示>pivot部分多了一个元素
			gt--
		} else {
			// 如果当前位置元素<pivot，则将当前位置元素与=pivot部分的第一个元素交换位置
			swap(nums, lt+1, i)
			//表示<pivot部分多了一个元素
			lt++
			//考虑下一个元素
			i++
		}
	}
	// 上面的遍历完成之后，将整个序列的"基准"元素放置到合适的位置,也就是将它放置在=pivot部分即可
	swap(nums, gt-1, high) //[如果枢纽元放在第一个位置 swap(nums,lt,low)]
	gt++                   //[如果枢纽元放在第一个位置 lt--]
	return lt, gt
}
