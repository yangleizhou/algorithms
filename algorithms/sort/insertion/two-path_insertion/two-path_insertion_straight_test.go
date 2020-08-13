package insertion

import (
	"fmt"
	"testing"
)

func TestTwoPathInsertionStraight(t *testing.T) {
	var nums = []int{59, 25, 100, 65, 78, 1, 5, 69, -1}
	fmt.Printf("in = %v\n", nums)
	fmt.Printf("out = %v\n", twoPathInsertionStraight(nums))
}
