package insertion

// twoPathInsertionStraight 2-路插入排序算法 & 直接插入排序
func twoPathInsertionStraight(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	first, last, n := 0, 0, len(nums)
	var temp = make([]int, n)
	temp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < temp[first] {
			first = (first - 1 + n) % n
			temp[first] = nums[i]
		} else if nums[i] >= temp[last] {
			last++
			temp[last] = nums[i]
		} else { //区间[first,last]插入nums[i] 使用直接插入排序算法
			last++
			k := last
			for ; nums[i] < temp[(k-1+n)%n]; k = (k - 1 + n) % n {
				temp[k] = temp[(k-1+n)%n]
			}
			temp[k] = nums[i]
		}
	}
	for i := 0; i < n; i++ {
		nums[i] = temp[(i+first)%n]
	}
	return nums
}
