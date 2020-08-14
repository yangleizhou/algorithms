package insertion

// twoPathInsertioBinary2   2-路插入排序算法-非循环数组 & 折半插入排序
func twoPathInsertioBinary2(nums []int) []int {
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
			low, high := fstart, fend
			for low <= high {
				mid := low + ((high - low) >> 1)
				if nums[mid] > x {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}
			for k := i; k > high+1; k-- {
				nums[k] = nums[k-1]
			}
			nums[high+1] = x
			fend++
			i++
		} else { //插入到后部有序表
			nums[i] = nums[j] //查看 2-path_insertion2.png
			low, high := rstart, rend
			for low <= high {
				mid := low + ((high - low) >> 1)
				if nums[mid] > x {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}
			for k := rstart; k < high+1; k++ { //移动元素
				nums[k-1] = nums[k]
			}
			nums[high] = x
			rstart--
			j--
		}
	}
	return nums
}
