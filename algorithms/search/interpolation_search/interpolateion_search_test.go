package search

import (
	"fmt"
	"testing"
)

func TestInterpolationSearch(t *testing.T) {
	var nums = []int{1, 2, 5, 7, 9, 23, 45, 67}
	var target int = 9

	index := interpolationSearch(nums, target)
	fmt.Println(index)
	if 4 != index {
		t.Errorf("index should be 4")
	}
	fmt.Println("")

	nums = []int{1, 2, 2, 2, 2, 2, 2, 2, 5, 5, 5, 5, 7, 9, 23, 45, 67}
	target = 2

	index = interpolationSearch(nums, target)
	fmt.Println(index)
	if 1 != index {
		t.Errorf("index should be 1")
	}
	fmt.Println("")

	nums = []int{1, 2}
	target = -1
	index = interpolationSearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 8}
	target = -1
	index = interpolationSearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 3}
	target = 9
	index = interpolationSearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 8}
	target = 9
	index = interpolationSearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}
}

func TestRecursiveInterpolationSearch(t *testing.T) {
	var nums = []int{1, 2, 5, 7, 9, 23, 45, 67}
	var target int = 9

	index := recursiveInterpolationSearch(nums, 0, len(nums)-1, target)
	fmt.Println(index)
	if 4 != index {
		t.Errorf("index should be 4")
	}

	nums = []int{1, 2, 2, 2, 2, 2, 2, 2, 5, 5, 5, 5, 7, 9, 23, 45, 67}
	target = 2

	index = recursiveInterpolationSearch(nums, 0, len(nums)-1, target)
	fmt.Println(index)
	if 1 != index {
		t.Errorf("index should be 1")
	}

	nums = []int{1, 2}
	target = -1
	index = recursiveInterpolationSearch(nums, 0, len(nums)-1, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 8}
	target = -1
	index = recursiveInterpolationSearch(nums, 0, len(nums)-1, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 3}
	target = 9
	index = recursiveInterpolationSearch(nums, 0, len(nums)-1, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 8}
	target = 9
	index = recursiveInterpolationSearch(nums, 0, len(nums)-1, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}
}
