package exchange

import (
	"fmt"
	"testing"
)

func TestQuickSortOptimizeInsertion(t *testing.T) { //快排优化 序列长度达到一定大小时，使用插入排序
	//nums := GenerateAscNums(30) //升序数组
	//nums := GenerateDesNums(20) //降序数组
	//nums := GenerateRepeatNums(10) //重复数组
	nums := GenerateRandNums(35) //随机数组
	fmt.Println(nums)
	quickSortOptimizeInsertion(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func TestQuickSortOptimizeTailRecursive(t *testing.T) { //快排优化  尾递归优化
	//nums := GenerateAscNums(30) //升序数组
	//nums := GenerateDesNums(20) //降序数组
	//nums := GenerateRepeatNums(10) //重复数组
	nums := GenerateRandNums(35) //随机数组
	fmt.Println(nums)
	quickSortOptimizeTailRecursive(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
