package search

import (
	"fmt"
	"testing"
)

func TestFibonacciSearch2(t *testing.T) {
	var nums = []int{1, 2, 5, 5, 7, 9, 23, 45, 67, 88, 98}
	var target int = 7

	index := fibonacciSearch2(nums, target)
	fmt.Println(index)

	if 4 != index {
		t.Errorf("index should be 4")
	}

	nums = []int{1, 2, 2, 2, 2, 2, 2, 5, 5, 5, 5, 5, 7, 7, 7, 7, 9, 23, 45, 67, 88}
	target = 45
	index = fibonacciSearch2(nums, target)
	fmt.Println(index)

	nums = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	target = 1
	index = fibonacciSearch2(nums, target)
	fmt.Println(index)

	nums = []int{1, 2}
	target = -1
	index = fibonacciSearch2(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 8}
	target = -1
	index = fibonacciSearch2(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 3}
	target = 9
	index = fibonacciSearch2(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 8}
	target = 9
	index = fibonacciSearch2(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}
}
