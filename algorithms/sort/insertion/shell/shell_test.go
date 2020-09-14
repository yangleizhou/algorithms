package insertion

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	var nums = []int{12, 2, 45, 8, 9, 23, 56, 45, 1}
	fmt.Println(nums)
	fmt.Println(shellSort(nums))

	nums = []int{59, 25, 100, 65, 78, 1, 5, 69, -1}
	fmt.Println(nums)
	fmt.Println(shellSort(nums))
}

func TestHibbardShellSort(t *testing.T) {
	var nums = []int{12, 2, 45, 8, 9, 23, 56, 45, 1}
	fmt.Println(nums)
	fmt.Println(hibbardShellSort(nums))

	nums = []int{59, 25, 100, 65, 78, 1, 5, 69, -1}
	fmt.Println(nums)
	fmt.Println(hibbardShellSort(nums))
}

func TestSedgewickShellSort(t *testing.T) {
	var nums = []int{12, 2, 45, 8, 9, 23, 56, 45, 1}
	fmt.Println(nums)
	fmt.Println(sedgewickShellSort(nums))

	nums = []int{59, 25, 100, 65, 78, 1, 5, 69, -1}
	fmt.Println(nums)
	fmt.Println(sedgewickShellSort(nums))
}
