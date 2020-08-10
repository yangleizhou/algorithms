package search

import (
	"fmt"
	"testing"
)

func TestFibonacciSearch(t *testing.T) {
	var nums = []int{1, 2, 5, 7, 9, 23, 45, 67, 88, 98}
	var target int = 9

	index := fibonacciSearch(nums, target)
	fmt.Println(index)
	if 4 != index {
		t.Errorf("index should be 4")
	}

	nums = []int{1, 2, 2, 2, 2, 2, 2, 5, 5, 5, 5, 5, 7, 7, 7, 7, 9, 23, 45, 67, 88}
	target = 2

	index = fibonacciSearch(nums, target)
	fmt.Println(index)

	nums = []int{1, 2}
	target = -1
	index = fibonacciSearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 8}
	target = -1
	index = fibonacciSearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 3}
	target = 9
	index = fibonacciSearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}

	nums = []int{1, 2, 8}
	target = 9
	index = fibonacciSearch(nums, target)
	fmt.Println(index)
	if -1 != index {
		t.Errorf("index should be -1")
	}
}
