package exchange

import (
	"math/rand"
	"time"
)

const (
	FixFirstPivot            = iota //固定基准first
	FixLastPivot                    //固定基准last
	RandPivot                       //随机基准
	MedianThreePivot                //三数取中
	MedianThreeOptimizePivot        //三数取中
)

// 获取基准
func getPivot(nums []int, low, high, typ int) int {
	switch typ {
	case FixFirstPivot:
		return FixedFirstPivot(nums, low, high)
	case FixLastPivot:
		return FixedLastPivot(nums, low, high)
	case RandPivot:
		return RandomPivot(nums, low, high)
	case MedianThreePivot:
		return MedianOfThreePivot(nums, low, high)
	case MedianThreeOptimizePivot:
		return MedianOfThreePivotOptimize(nums, low, high)
	}
	return -1
}

// 交换两数
func swap(nums []int, low, high int) {
	nums[low], nums[high] = nums[high], nums[low]
}

// FixedFirstPivot 固定基准-第一个
func FixedFirstPivot(nums []int, low, high int) int {
	swap(nums, low, high)
	return nums[high]
}

// FixedLastPivot 固定基准-最后
func FixedLastPivot(nums []int, low, high int) int {
	return nums[high]
}

// RandomPivot 随机基准
func RandomPivot(nums []int, low, high int) int {
	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Int()%(high-low) + low
	swap(nums, high, pivotIndex)
	return nums[high]
}

func MedianOfThreePivot(nums []int, low, high int) int {
	mid := low + ((high - low) >> 1)
	//找出三个数中的最小值放到 nums[low]
	if nums[mid] < nums[low] {
		swap(nums, low, mid)
	}
	if nums[high] < nums[low] {
		swap(nums, low, high)
	}
	//将 中间那个数放到 nums[media]
	if nums[mid] > nums[high] {
		swap(nums, mid, high)
	}
	//尽量将大的元素放到右边--将privot放到右边, 可简化 分割操作(partition)
	swap(nums, mid, high) //这里使用high-1 不是high 因为通过上面比较nums[high]一定大于等于nums[mid]
	return nums[high]
}

// MedianOfThreePivotOptimize 三数取中优化
func MedianOfThreePivotOptimize(nums []int, low, high int) int {
	mid := low + ((high - low) >> 1)
	//找出三个数中的最小值放到 nums[low]
	if nums[mid] < nums[low] {
		swap(nums, low, mid)
	}
	if nums[high] < nums[low] {
		swap(nums, low, high)
	}
	//将 中间那个数放到 nums[media]
	if nums[mid] > nums[high] {
		swap(nums, mid, high)
	}
	//尽量将大的元素放到右边--将privot放到右边, 可简化 分割操作(partition)
	swap(nums, mid, high-1) //这里使用high-1 不是high 因为通过上面比较nums[high]一定大于等于nums[mid]
	return nums[high-1]
}
