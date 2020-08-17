package insertion

import (
	"fmt"
	"testing"
)

func TestTwoPathInsertionStraight2(t *testing.T) {

	var nums = []int{7, 5, 45, 78, 65, 59, 25, 100, 65, 78, 1, 5, 69, 6}
	fmt.Printf("in = %v\n", nums)
	fmt.Printf("out = %v\n", twoPathInsertionStraight2(nums))

	nums = []int{59, 25, 100, 65, 78, 1, 5, 69, -1}
	fmt.Printf("in = %v\n", nums)
	fmt.Printf("out = %v\n", twoPathInsertionStraight2(nums))

}
