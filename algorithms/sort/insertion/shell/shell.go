package insertion

// 原始希尔排序
// 增量序列为D(M)=N/2, D(k)=D(k+1)/2 向下取整
func shellSort(nums []int) []int {
	size := len(nums)
	step := size >> 1
	for step > 0 {
		for i := step; i < size; i += step {
			j := i
			temp := nums[j]
			for j >= step {
				if temp >= nums[j-step] {
					break
				}
				nums[j] = nums[j-step]
				j -= step
			}
			nums[j] = temp
		}
		step = step >> 1
	}
	return nums
}

// 希尔排序（Hibbard增量序列）
// 增量序列 D(k)=2^k−1
func hibbardShellSort(nums []int) []int {
	size := len(nums)
	steps := getHibbardSteps(size)
	for k := len(steps) - 1; k >= 0; k-- {
		step := steps[k]
		for i := step; i < size; i += step {
			j := i
			temp := nums[j]
			for j >= step {
				if temp >= nums[j-step] {
					break
				}
				nums[j] = nums[j-step]
				j -= step
			}
			nums[j] = temp
		}
	}
	return nums
}

// Hibbard增量序列
// D(i)=2^i−1,i>0
func getHibbardSteps(size int) (steps []int) {
	i := 1
	temp := (1 << i) - 1
	for temp <= size {
		steps = append(steps, temp)
		i++
		temp = (1 << i) - 1
	}
	return
}

// 希尔排序（Sedgewick增量序列）
func sedgewickShellSort(nums []int) []int {
	size := len(nums)
	steps := getSedgewichSteps(size)
	for k := len(steps) - 1; k >= 0; k-- {
		step := steps[k]
		for i := step; i < size; i += step {
			j := i
			temp := nums[j]
			for j >= step {
				if temp >= nums[j-step] {
					break
				}
				nums[j] = nums[j-step]
				j -= step
			}
			nums[j] = temp
		}
	}
	return nums
}

// Sedgewick增量序列
// D=9*4^i-9*2^i+1 或 4^(i+2)-3*2^(i+2)+1 , i>=0
// 稍微变一下形：D=9*(2^(2i)-2^i)+1 或 2^(2i+4)-3*2^(i+2)+1 , i>=0
func getSedgewichSteps(size int) (steps []int) {
	i := 0
	for {
		temp := 9*((1<<(2*i))-(1<<i)) + 1
		if temp <= size {
			steps = append(steps, temp)
		}
		temp = (1 << (2*i + 4)) - 3*(1<<(i+2)) + 1
		if temp <= size {
			steps = append(steps, temp)
		} else {
			break
		}
		i++
	}
	return

}
