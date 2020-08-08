package insertion

import (
	"fmt"
	"testing"
)

func TestStraightInsertion(t *testing.T) {
	var nums = []int{20, 3, 5, 2, 89, 6, 2, 4, 3, 10, 11, 7, 80}
	fmt.Println("input:\n", nums)
	fmt.Println("value:\n", straightInsertion(nums))

}
