package binary

import (
	"fmt"
	"testing"
)

func TestCountTriplets(t *testing.T) {
	var arr = []int{2, 3, 1, 6, 7}
	countTriplets(arr)

	fmt.Println()
	arr = []int{1, 1, 1, 1, 1}
	countTriplets(arr)

	fmt.Println()
	arr = []int{2, 3}
	countTriplets(arr)

	fmt.Println()
	arr = []int{1, 3, 5, 7, 9}
	countTriplets(arr)

	fmt.Println()
	arr = []int{7, 11, 12, 9, 5, 2, 7, 17, 22}
	countTriplets(arr)

}
