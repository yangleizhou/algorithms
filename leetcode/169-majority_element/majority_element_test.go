package binary

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	var nums = []int{3, 2, 3}
	value := majorityElement(nums)
	fmt.Println(value)

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	majorityElement(nums)
	fmt.Println(value)
}
