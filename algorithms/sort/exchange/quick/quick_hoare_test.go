package exchange

import (
	"fmt"
	"testing"
)

func TestQuickSortHoare(t *testing.T) {
	//nums := GenerateAscNums(10) //升序数组
	//nums := GenerateDesNums(10) //降序数组
	//nums := GenerateRepeatNums(10) //重复数组
	nums := GenerateRandNums(26) //随机数组
	fmt.Println(nums)
	quickSortHoare(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func TestQuickSortMedianOptimizeHoare(t *testing.T) {
	//nums := GenerateAscNums(10) //升序数组
	//nums := GenerateDesNums(10) //降序数组
	//nums := GenerateRepeatNums(10) //重复数组
	nums := GenerateRandNums(25) //随机数组
	fmt.Println(nums)
	quickSortMedianOptimizeHoare(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
