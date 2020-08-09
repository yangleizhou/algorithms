package search

import (
	"fmt"
	"testing"
)

func TestRightBoundBinarySearch(t *testing.T) {
	var nums = []int{1, 2, 2, 2, 4, 5}
	var target int = 2
	index := rightBoundBinarySearch(nums, target)
	fmt.Println(index)
	if 3 != index {
		t.Errorf("index should be 3")
	}

	nums = []int{1, 3, 5, 9, 34, 45}
	target = 9
	index = rightBoundBinarySearch(nums, target)
	fmt.Println(index)
	if 3 != index {
		t.Errorf("index should be 3")
	}

	nums = []int{1, 3, 5, 9, 34, 45}
	target = 10
	index = rightBoundBinarySearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 3, 5, 9, 34, 45}
	target = 60
	index = rightBoundBinarySearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}
}
