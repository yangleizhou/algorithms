package insertion

// twoPathInsertion2  2-路插入排序算法-非循环数组 & 直接插入排序
func twoPathInsertion2(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	n := len(nums)
	fstart, fend, rstart, rend := 0, 0, n-1, n-1
	if nums[0] > nums[n-1] { //形成前后部分只有一个元素有序的列表
		nums[0] = nums[0] ^ nums[n-1]
		nums[n-1] = nums[0] ^ nums[n-1]
		nums[0] = nums[0] ^ nums[n-1]
	}
	i, j := 1, n-2
	for i <= j {
		x := nums[i]
		if x < nums[rstart] { //插入到前部有序表
			k := fend
			for ; k >= fstart && x < nums[k]; k-- { //中间判断 k>=fstart 一定要在前面,否则下次k--后,再次进入判断nums[k]导致越界
				nums[k+1] = nums[k]
			}
			nums[k+1] = x
			fend++
			i++
		} else { //插入到后部有序表
			nums[i] = nums[j] //查看 2-path_insertion2.png
			k := rstart
			for ; k <= rend && x > nums[k]; k++ {
				nums[k-1] = nums[k]
			}
			nums[k-1] = x
			rstart--
			j--
		}
	}
	return nums
}
