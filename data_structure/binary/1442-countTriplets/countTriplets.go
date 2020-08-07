package binary

import "fmt"

func countTriplets(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	var count int
	for i := 0; i <= len(arr)-1; i++ {
		var temp = arr[i]
		for j := i + 1; j <= len(arr)-1; j++ {
			temp ^= arr[j]
			if (temp) == 0 {
				count += (j - i)
				for k := i + 1; k <= j; k++ {
					fmt.Printf("(%d,%d,%d)\n", i, k, j)
				}
			}
		}
	}
	fmt.Println(count)
	return count
}
