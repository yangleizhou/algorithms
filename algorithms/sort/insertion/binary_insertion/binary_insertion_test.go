package insertion

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBinaryInsertion(t *testing.T) {
	var nums = []int{12, 2, 45, 8, 9, 23, 56, 45, 1}
	fmt.Println(nums)
	fmt.Println(binaryInsertion(nums))

	nums = []int{59, 25, 100, 65, 78, 1, 5, 69, -1, 8, 9, 4}
	fmt.Println(nums)
	fmt.Println(binaryInsertion(nums))

}

func TestBinaryInsertion1(t *testing.T) {
	var nums = []int{12, 2, 45, 8, 9, 23, 56, 45, 1}
	fmt.Println(nums)
	BinaryInsertion(nums, 0, len(nums)-1)
	fmt.Println(nums)
	fmt.Println()

	nums = []int{59, 25, 100, 65, 78, 1, 5, 69, -1, 8, 9, 4}
	fmt.Println(nums)
	BinaryInsertion(nums, 0, len(nums)-1)
	fmt.Println(nums)
	fmt.Println()

	nums = []int{10, 15, 11, 14, 15}
	fmt.Println(nums)
	BinaryInsertion(nums, 0, len(nums)-1)
	fmt.Println(nums)
	fmt.Println()

	size := 30
	nums = make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		nums[i] = rand.Int() % size
	}
	fmt.Println(nums)
	BinaryInsertion(nums, 0, len(nums)-1)
	fmt.Println(nums)
	fmt.Println()

}
