package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var nums = []int{1, 2, 5, 7, 9, 23, 45, 67}
	var target int = 9

	index := binarySearch(nums, target)
	fmt.Println(index)
	if 4 != index {
		t.Errorf("index should be 4")
	}

	nums = []int{1, 2}
	target = -1
	index = binarySearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 8}
	target = -1
	index = binarySearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 3}
	target = 9
	index = binarySearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 8}
	target = 9
	index = binarySearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}
}

func TestRecursiveBinarySearch(t *testing.T) {
	var nums = []int{1, 2, 5, 7, 9, 23, 45, 67}
	var target int = 9

	index := recursiveBinarySearch(nums, 0, len(nums)-1, target)
	fmt.Println(index)
	if 4 != index {
		t.Errorf("index should be 4")
	}

	nums = []int{1, 2}
	target = -1
	index = recursiveBinarySearch(nums, 0, len(nums)-1, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 8}
	target = -1
	index = recursiveBinarySearch(nums, 0, len(nums)-1, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 3}
	target = 9
	index = recursiveBinarySearch(nums, 0, len(nums)-1, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 8}
	target = 9
	index = recursiveBinarySearch(nums, 0, len(nums)-1, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}
}
