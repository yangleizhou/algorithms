package insertion

import (
	"fmt"
	"testing"
)

func TestBinaryInsertion(t *testing.T) {
	var nums = []int{12, 2, 45, 8, 9, 23, 56, 45, 1}
	fmt.Println(nums)
	fmt.Println(binaryInsertion(nums))
}
