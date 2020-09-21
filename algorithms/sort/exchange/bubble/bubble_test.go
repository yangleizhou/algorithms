package exchange

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	var nums = []int{12, 2, 45, 8, 9, 23, 56, 45, 1}
	fmt.Println(nums)
	fmt.Println(bubble(nums))

	nums = []int{12, 2, 45, 8, 9, 23, 56, 45, 1}
	recursiveBubble(nums, len(nums))
	fmt.Println(nums)
}
