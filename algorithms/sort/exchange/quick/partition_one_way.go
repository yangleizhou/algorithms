package exchange

//单向扫描

func oneWayPartition(nums []int, low, high, pivotType int) int {
	pivot := getPivot(nums, low, high, pivotType)
	m, i := low, low         //如果枢纽放在第一个位置上  m, i := low, low+1
	for ; i <= high-1; i++ { // 如果枢纽放在第一个位置上 i<=high
		if nums[i] < pivot {
			swap(nums, m, i)
			m++
		}
	}
	swap(nums, m, high) //如果枢纽放在第一个位置上 swap(nums, m, low)
	return m
}
