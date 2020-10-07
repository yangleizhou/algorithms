package exchange

import (
	"fmt"
	"testing"
)

func TestQuickSortThreeWay(t *testing.T) {
	//nums := GenerateAscNums(10) //升序数组
	//nums := GenerateDesNums(10) //降序数组
	//nums := GenerateRepeatNums(10) //重复数组
	nums := GenerateRandNums(34) //随机数组
	fmt.Println(nums)
	quickSortThreeWay(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
