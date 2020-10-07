package exchange

import (
	"fmt"
	"testing"
)

func TestQuickSortOneWay(t *testing.T) {
	nums := GenerateAscNums(30) //升序数组
	//nums := GenerateDesNums(20) //降序数组
	//nums := GenerateRepeatNums(10) //重复数组
	//nums := GenerateRandNums(30) //随机数组
	fmt.Println(nums)
	quickSortOneWay(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
