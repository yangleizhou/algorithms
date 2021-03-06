package insertion

import (
	"fmt"
	"testing"
)

func TestTwoPathInsertionBinary(t *testing.T) {
	var nums = []int{59, 25, 100, 65, 78, 1, 5, 69, -1}
	fmt.Printf("in = %v\n", nums)
	fmt.Printf("out = %v\n", twoPathInsertionBinary(nums))

	nums = []int{7, 23, 45, 78, 65, 59, 25, 100, 65, 78, 1, 5, 69, -1}
	fmt.Printf("in = %v\n", nums)
	fmt.Printf("out = %v\n", twoPathInsertionBinary(nums))
}
