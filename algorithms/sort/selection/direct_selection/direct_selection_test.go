package selection

import (
	"fmt"
	"testing"
)

func TestDirectSelection(t *testing.T) {
	var nums = []int{12, 2, 45, 8, 9, 23, 56, 45, 1}
	fmt.Println(nums)
	fmt.Println(directSelection(nums))

	nums = []int{59, 25, 100, 65, 78, 1, 5, 69, -1, 8, 9, 4}
	fmt.Println(nums)
	fmt.Println(directSelection(nums))
}
