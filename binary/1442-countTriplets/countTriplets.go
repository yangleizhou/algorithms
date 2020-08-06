package binary

import "fmt"

func countTriplets(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	var count int
	for i := len(arr) - 1; i > 0; i-- {
		var temp = arr[i]
		for j := i - 1; j > 0; j-- {
			temp ^= arr[j]
			if temp == 0 {
				count += (j - i - 1)
				for k := j + 1; k < j; k++ {
					fmt.Printf("(%d,%d,%d)\n", j, k, i)
				}
			}
		}
	}
	return count
}
